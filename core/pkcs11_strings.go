package core

import "fmt"

func CKAString(attr uint32) string {
	switch attr {
	case CKA_CLASS:
		return "CKA_CLASS"
	case CKA_TOKEN:
		return "CKA_TOKEN"
	case CKA_PRIVATE:
		return "CKA_PRIVATE"
	case CKA_LABEL:
		return "CKA_LABEL"
	case CKA_APPLICATION:
		return "CKA_APPLICATION"
	case CKA_VALUE:
		return "CKA_VALUE"
	case CKA_OBJECT_ID:
		return "CKA_OBJECT_ID"
	case CKA_CERTIFICATE_TYPE:
		return "CKA_CERTIFICATE_TYPE"
	case CKA_ISSUER:
		return "CKA_ISSUER"
	case CKA_SERIAL_NUMBER:
		return "CKA_SERIAL_NUMBER"
	case CKA_AC_ISSUER:
		return "CKA_AC_ISSUER"
	case CKA_OWNER:
		return "CKA_OWNER"
	case CKA_ATTR_TYPES:
		return "CKA_ATTR_TYPES"
	case CKA_TRUSTED:
		return "CKA_TRUSTED"
	case CKA_CERTIFICATE_CATEGORY:
		return "CKA_CERTIFICATE_CATEGORY"
	case CKA_JAVA_MIDP_SECURITY_DOMAIN:
		return "CKA_JAVA_MIDP_SECURITY_DOMAIN"
	case CKA_URL:
		return "CKA_URL"
	case CKA_HASH_OF_SUBJECT_PUBLIC_KEY:
		return "CKA_HASH_OF_SUBJECT_PUBLIC_KEY"
	case CKA_HASH_OF_ISSUER_PUBLIC_KEY:
		return "CKA_HASH_OF_ISSUER_PUBLIC_KEY"
	case CKA_NAME_HASH_ALGORITHM:
		return "CKA_NAME_HASH_ALGORITHM"
	case CKA_CHECK_VALUE:
		return "CKA_CHECK_VALUE"
	case CKA_KEY_TYPE:
		return "CKA_KEY_TYPE"
	case CKA_SUBJECT:
		return "CKA_SUBJECT"
	case CKA_ID:
		return "CKA_ID"
	case CKA_SENSITIVE:
		return "CKA_SENSITIVE"
	case CKA_ENCRYPT:
		return "CKA_ENCRYPT"
	case CKA_DECRYPT:
		return "CKA_DECRYPT"
	case CKA_WRAP:
		return "CKA_WRAP"
	case CKA_UNWRAP:
		return "CKA_UNWRAP"
	case CKA_SIGN:
		return "CKA_SIGN"
	case CKA_SIGN_RECOVER:
		return "CKA_SIGN_RECOVER"
	case CKA_VERIFY:
		return "CKA_VERIFY"
	case CKA_VERIFY_RECOVER:
		return "CKA_VERIFY_RECOVER"
	case CKA_DERIVE:
		return "CKA_DERIVE"
	case CKA_START_DATE:
		return "CKA_START_DATE"
	case CKA_END_DATE:
		return "CKA_END_DATE"
	case CKA_MODULUS:
		return "CKA_MODULUS"
	case CKA_MODULUS_BITS:
		return "CKA_MODULUS_BITS"
	case CKA_PUBLIC_EXPONENT:
		return "CKA_PUBLIC_EXPONENT"
	case CKA_PRIVATE_EXPONENT:
		return "CKA_PRIVATE_EXPONENT"
	case CKA_PRIME_1:
		return "CKA_PRIME_1"
	case CKA_PRIME_2:
		return "CKA_PRIME_2"
	case CKA_EXPONENT_1:
		return "CKA_EXPONENT_1"
	case CKA_EXPONENT_2:
		return "CKA_EXPONENT_2"
	case CKA_COEFFICIENT:
		return "CKA_COEFFICIENT"
	case CKA_PUBLIC_KEY_INFO:
		return "CKA_PUBLIC_KEY_INFO"
	case CKA_PRIME:
		return "CKA_PRIME"
	case CKA_SUBPRIME:
		return "CKA_SUBPRIME"
	case CKA_BASE:
		return "CKA_BASE"
	case CKA_PRIME_BITS:
		return "CKA_PRIME_BITS"
	case CKA_SUBPRIME_BITS:
		return "CKA_SUBPRIME_BITS"
	case CKA_VALUE_BITS:
		return "CKA_VALUE_BITS"
	case CKA_VALUE_LEN:
		return "CKA_VALUE_LEN"
	case CKA_EXTRACTABLE:
		return "CKA_EXTRACTABLE"
	case CKA_LOCAL:
		return "CKA_LOCAL"
	case CKA_NEVER_EXTRACTABLE:
		return "CKA_NEVER_EXTRACTABLE"
	case CKA_ALWAYS_SENSITIVE:
		return "CKA_ALWAYS_SENSITIVE"
	case CKA_KEY_GEN_MECHANISM:
		return "CKA_KEY_GEN_MECHANISM"
	case CKA_MODIFIABLE:
		return "CKA_MODIFIABLE"
	case CKA_COPYABLE:
		return "CKA_COPYABLE"
	case CKA_DESTROYABLE:
		return "CKA_DESTROYABLE"
	case CKA_EC_PARAMS:
		return "CKA_EC_PARAMS"
	case CKA_EC_POINT:
		return "CKA_EC_POINT"
	case CKA_SECONDARY_AUTH:
		return "CKA_SECONDARY_AUTH"
	case CKA_AUTH_PIN_FLAGS:
		return "CKA_AUTH_PIN_FLAGS"
	case CKA_ALWAYS_AUTHENTICATE:
		return "CKA_ALWAYS_AUTHENTICATE"
	case CKA_WRAP_WITH_TRUSTED:
		return "CKA_WRAP_WITH_TRUSTED"
	case CKA_WRAP_TEMPLATE:
		return "CKA_WRAP_TEMPLATE"
	case CKA_UNWRAP_TEMPLATE:
		return "CKA_UNWRAP_TEMPLATE"
	// case C.CKA_DERIVE_TEMPLATE:
	// 	return "CKA_DERIVE_TEMPLATE"
	case CKA_OTP_FORMAT:
		return "CKA_OTP_FORMAT"
	case CKA_OTP_LENGTH:
		return "CKA_OTP_LENGTH"
	case CKA_OTP_TIME_INTERVAL:
		return "CKA_OTP_TIME_INTERVAL"
	case CKA_OTP_USER_FRIENDLY_MODE:
		return "CKA_OTP_USER_FRIENDLY_MODE"
	case CKA_OTP_CHALLENGE_REQUIREMENT:
		return "CKA_OTP_CHALLENGE_REQUIREMENT"
	case CKA_OTP_TIME_REQUIREMENT:
		return "CKA_OTP_TIME_REQUIREMENT"
	case CKA_OTP_COUNTER_REQUIREMENT:
		return "CKA_OTP_COUNTER_REQUIREMENT"
	case CKA_OTP_PIN_REQUIREMENT:
		return "CKA_OTP_PIN_REQUIREMENT"
	case CKA_OTP_COUNTER:
		return "CKA_OTP_COUNTER"
	case CKA_OTP_TIME:
		return "CKA_OTP_TIME"
	case CKA_OTP_USER_IDENTIFIER:
		return "CKA_OTP_USER_IDENTIFIER"
	case CKA_OTP_SERVICE_IDENTIFIER:
		return "CKA_OTP_SERVICE_IDENTIFIER"
	case CKA_OTP_SERVICE_LOGO:
		return "CKA_OTP_SERVICE_LOGO"
	case CKA_OTP_SERVICE_LOGO_TYPE:
		return "CKA_OTP_SERVICE_LOGO_TYPE"
	case CKA_GOSTR3410_PARAMS:
		return "CKA_GOSTR3410_PARAMS"
	case CKA_GOSTR3411_PARAMS:
		return "CKA_GOSTR3411_PARAMS"
	case CKA_GOST28147_PARAMS:
		return "CKA_GOST28147_PARAMS"
	case CKA_HW_FEATURE_TYPE:
		return "CKA_HW_FEATURE_TYPE"
	case CKA_RESET_ON_INIT:
		return "CKA_RESET_ON_INIT"
	case CKA_HAS_RESET:
		return "CKA_HAS_RESET"
	case CKA_PIXEL_X:
		return "CKA_PIXEL_X"
	case CKA_PIXEL_Y:
		return "CKA_PIXEL_Y"
	case CKA_RESOLUTION:
		return "CKA_RESOLUTION"
	case CKA_CHAR_ROWS:
		return "CKA_CHAR_ROWS"
	case CKA_CHAR_COLUMNS:
		return "CKA_CHAR_COLUMNS"
	case CKA_COLOR:
		return "CKA_COLOR"
	case CKA_BITS_PER_PIXEL:
		return "CKA_BITS_PER_PIXEL"
	case CKA_CHAR_SETS:
		return "CKA_CHAR_SETS"
	case CKA_ENCODING_METHODS:
		return "CKA_ENCODING_METHODS"
	case CKA_MIME_TYPES:
		return "CKA_MIME_TYPES"
	case CKA_MECHANISM_TYPE:
		return "CKA_MECHANISM_TYPE"
	case CKA_REQUIRED_CMS_ATTRIBUTES:
		return "CKA_REQUIRED_CMS_ATTRIBUTES"
	case CKA_DEFAULT_CMS_ATTRIBUTES:
		return "CKA_DEFAULT_CMS_ATTRIBUTES"
	case CKA_SUPPORTED_CMS_ATTRIBUTES:
		return "CKA_SUPPORTED_CMS_ATTRIBUTES"
	case CKA_ALLOWED_MECHANISMS:
		return "CKA_ALLOWED_MECHANISMS"
	case CKA_VENDOR_DEFINED:
		return "CKA_VENDOR_DEFINED"
	default:
		return fmt.Sprint(attr)
	}
}
