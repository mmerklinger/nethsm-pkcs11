use merge::Merge;
use serde::{Deserialize, Serialize};

#[derive(Debug)]
pub enum ConfigError {
    Io(std::io::Error),
    Yaml(serde_yaml::Error),
    NoConfigFile,
}

const CONFIG_FILE_NAME: &str = "p11nethsm.conf";
const ENV_VAR_CONFIG_FILE: &str = "P11NETHSM_CONFIG_FILE";

pub fn read_configuration() -> Result<P11Config, ConfigError> {
    let mut config = P11Config::default();

    if let Ok(file_path) = std::env::var(ENV_VAR_CONFIG_FILE) {
        let file = std::fs::File::open(file_path).map_err(ConfigError::Io)?;
        let config_file = serde_yaml::from_reader(file).map_err(ConfigError::Yaml)?;

        config.merge(config_file);

        return Ok(config);
    }

    let mut config_folders = vec![
        "/etc/nitrokey".to_string(),
        "/usr/local/etc/nitrokey".to_string(),
    ];

    if let Ok(home) = std::env::var("HOME") {
        config_folders.push(format!("{}/.config/nitrokey", home));
    }

    let mut file_read = false;

    for folder in config_folders {
        let file_path = format!("{}/{}", folder, CONFIG_FILE_NAME);

        if let Ok(file) = std::fs::File::open(file_path) {
            let config_file = serde_yaml::from_reader(file).map_err(ConfigError::Yaml)?;

            config.merge(config_file);
            file_read = true;
        }
    }

    // if no config file was found, return an error

    if !file_read {
        return Err(ConfigError::NoConfigFile);
    }

    Ok(config)
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub enum LogLevel {
    Trace,
    Debug,
    Info,
    Warn,
    Error,
}

impl From<&LogLevel> for log::LevelFilter {
    fn from(level: &LogLevel) -> Self {
        match level {
            LogLevel::Trace => log::LevelFilter::Trace,
            LogLevel::Debug => log::LevelFilter::Debug,
            LogLevel::Info => log::LevelFilter::Info,
            LogLevel::Warn => log::LevelFilter::Warn,
            LogLevel::Error => log::LevelFilter::Error,
        }
    }
}

// representation of the config file to parse
#[derive(Debug, Clone, Serialize, Deserialize, Merge, Default)]
pub struct P11Config {
    #[merge(strategy = merge::bool::overwrite_false)]
    #[serde(default)]
    pub enable_set_attribute_value: bool,
    pub log_file: Option<String>,
    pub log_level: Option<LogLevel>,
    #[merge(strategy = merge::vec::append)]
    pub slots: Vec<SlotConfig>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct InstanceConfig {
    pub url: String,
    #[serde(default)]
    pub danger_insecure_cert: bool,
    #[serde(default)]
    pub sha256_fingerprints: Vec<String>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SlotConfig {
    pub label: String,
    pub operator: Option<UserConfig>,
    pub administrator: Option<UserConfig>,
    pub description: Option<String>,
    pub instances: Vec<InstanceConfig>,
}

// An user
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UserConfig {
    pub username: String,
    #[serde(deserialize_with = "deserialize_password", default)]
    pub password: Option<String>,
}

const PASSWORD_ENV_PREFIX: &str = "env:";

// Deserialize a string, but if it starts with "env:" then read the environment variable corresponding to the rest of the string
fn deserialize_password<'de, D>(deserializer: D) -> Result<Option<String>, D::Error>
where
    D: serde::Deserializer<'de>,
{
    match Option::<String>::deserialize(deserializer)? {
        Some(s) => {
            if s.starts_with(PASSWORD_ENV_PREFIX) {
                let var = s.trim_start_matches(PASSWORD_ENV_PREFIX);
                let val = std::env::var(var).map_err(serde::de::Error::custom)?;
                return Ok(Some(val));
            }
            if s.is_empty() {
                return Ok(None);
            }
            Ok(Some(s))
        }
        None => Ok(None),
    }
}

#[cfg(test)]
mod tests {
    use std::fs;

    use super::*;

    #[test]
    fn test_read_home_config() {
        let config = r#"
enable_set_attribute_value: true
log_file: /tmp/p11nethsm.log
log_level: Trace
slots:
  - label: test
    operator:
        username: test
        password: test_password
    instances:
        - url: https://localhost:23443
          danger_insecure_cert: true
"#;
        let home = "/tmp/home/";

        // create a temporary "fake" home folder
        fs::create_dir_all(format!("{}.config/nitrokey", home)).unwrap();
        fs::write(
            format!("{}/.config/nitrokey/{}", home, CONFIG_FILE_NAME),
            config,
        )
        .unwrap();

        std::env::remove_var(ENV_VAR_CONFIG_FILE);

        std::env::set_var("HOME", home);

        let config = read_configuration().unwrap();
        assert!(config.enable_set_attribute_value);
        assert_eq!(config.log_file, Some("/tmp/p11nethsm.log".to_string()));
        assert!(matches!(config.log_level, Some(LogLevel::Trace)));
        assert_eq!(config.slots.len(), 1);
        assert_eq!(config.slots[0].label, "test");
        assert_eq!(config.slots[0].operator.as_ref().unwrap().username, "test");
        assert_eq!(
            config.slots[0].operator.as_ref().unwrap().password,
            Some("test_password".to_string())
        );
        // clean up
        fs::remove_dir_all(home).unwrap();
    }

    #[test]
    #[ignore]
    fn test_read_config_no_file() {
        std::env::remove_var(ENV_VAR_CONFIG_FILE);
        let config = read_configuration();
        assert!(config.is_err());
        assert!(matches!(
            config.unwrap_err(),
            super::ConfigError::NoConfigFile
        ));
    }

    #[test]
    fn test_deserialize_password_env() {
        let config = r#"
username: test
password: env:TEST_PASSWORD
"#;

        std::env::set_var("TEST_PASSWORD", "test_password");
        let config: super::UserConfig = serde_yaml::from_str(config).unwrap();
        assert_eq!(config.username, "test");
        assert_eq!(config.password, Some("test_password".to_string()));
    }

    #[test]
    fn test_deserialize_password() {
        let config = r#"
username: test
password: test_password
"#;
        let config: super::UserConfig = serde_yaml::from_str(config).unwrap();
        assert_eq!(config.username, "test");
        assert_eq!(config.password, Some("test_password".to_string()));
    }

    #[test]
    fn test_deserialize_password_none() {
        let config = r#"
username: test
"#;
        let config: super::UserConfig = serde_yaml::from_str(config).unwrap();
        assert_eq!(config.username, "test");
        assert_eq!(config.password, None);
    }

    #[test]
    fn test_deserialize_password_empty() {
        let config = r#"
username: test
password: ""
"#;
        let config: super::UserConfig = serde_yaml::from_str(config).unwrap();
        assert_eq!(config.username, "test");
        assert_eq!(config.password, None);
    }
}
