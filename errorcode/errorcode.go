package errorcode

type CodeDescription struct {
	HttpStatusCode int
	Code           string
	Message        string
}

var SUCCESS = CodeDescription{
	HttpStatusCode: 200,
	Code:           "SUCCESS",
	Message:        "성공.",
}

var CREATED = CodeDescription{
	HttpStatusCode: 201,
	Code:           "CREATED",
	Message:        "생성 성공.",
}

var NO_CONTENT = CodeDescription{
	HttpStatusCode: 204,
	Code:           "NO_CONTENT",
	Message:        "전송할 데이터가 없습니다.",
}

var TOKEN_MISSING = CodeDescription{
	HttpStatusCode: 400,
	Code:           "TOKEN_MISSING",
	Message:        "토큰을 찾을 수 없습니다.",
}

var TOKEN_INVALID_FORMAT = CodeDescription{
	HttpStatusCode: 400,
	Code:           "TOKEN_INVALID_FORMAT",
	Message:        "올바른 형식이 아닌 토큰입니다.",
}

var TOKEN_PARSING_ERROR = CodeDescription{
	HttpStatusCode: 400,
	Code:           "TOKEN_PARSING_ERROR",
	Message:        "올바른 형식이 아닌 토큰입니다.",
}

var TOKEN_VERIFICATION_ERROR = CodeDescription{
	HttpStatusCode: 400,
	Code:           "TOKEN_VERIFICATION_ERROR",
	Message:        "토큰 검증을 실패했습니다.",
}

var TOKEN_EXPIRED = CodeDescription{
	HttpStatusCode: 400,
	Code:           "TOKEN_EXPIRED",
	Message:        "토큰이 만료되었습니다.",
}

var REFRESH_TOKEN_EXPIRED = CodeDescription{
	HttpStatusCode: 400,
	Code:           "REFRESH_TOKEN_EXPIRED",
	Message:        "토큰이 만료되었습니다.",
}

var BAD_REQUEST = CodeDescription{
	HttpStatusCode: 400,
	Code:           "BAD_REQUEST",
	Message:        "잘못된 요청입니다.▸ ",
}

var MISSING_PARAMETERS = CodeDescription{
	HttpStatusCode: 400,
	Code:           "MISSING_PARAMETERS",
	Message:        "필수 입력 정보가 부족합니다.▸ ",
}

var INVALID_PARAMETERS = CodeDescription{
	HttpStatusCode: 400,
	Code:           "INVALID_PARAMETERS",
	Message:        "입력 정보가 올바르지 않습니다.▸ ",
}

var INVALID_OTP = CodeDescription{
	HttpStatusCode: 400,
	Code:           "INVALID_OTP",
	Message:        "OTP 번호가 잘못되었거나 유효만료 시간을 초과했습니다.",
}

var ACCESS_DENIED = CodeDescription{
	HttpStatusCode: 401,
	Code:           "ACCESS_DENIED",
	Message:        "잘못된 IAM 키를 사용하여 접근했습니다. ",
}

var ACCESS_DENIED_ACCOUNT_DISABLE = CodeDescription{
	HttpStatusCode: 401,
	Code:           "ACCESS_DENIED_ACCOUNT_DISABLE",
	Message:        "계정 잠금 조치(비활성화)가 진행되었습니다. ",
}

var RESOLVE_REQUIRED_ACTIONS = CodeDescription{
	HttpStatusCode: 401,
	Code:           "RESOLVE_REQUIRED_ACTIONS",
	Message:        "비밀번호를 재설정해주세요. 90일이 만료되었습니다.",
}

var AUTH_EMAIL_ALREADY_EXISTS = CodeDescription{
	HttpStatusCode: 401,
	Code:           "AUTH_EMAIL_ALREADY_EXISTS",
	Message:        "이미 다른 계정에서 동일한 이메일을 사용하고 있습니다.▸ ",
}

var FORBIDDEN_REQUEST = CodeDescription{
	HttpStatusCode: 403,
	Code:           "FORBIDDEN_REQUEST",
	Message:        "허용되지 않은 요청입니다. ",
}

var NOT_FOUND_ERROR = CodeDescription{
	HttpStatusCode: 404,
	Code:           "NOT_FOUND_ERROR",
	Message:        "존재하지 않는 정보입니다.▸ ",
}

var DUPLICATED_KEY = CodeDescription{
	HttpStatusCode: 409,
	Code:           "DUPLICATED_KEY",
	Message:        "중복된 요청입니다.",
}

var TOO_MANY_REQUEST = CodeDescription{
	HttpStatusCode: 429,
	Code:           "TOO_MANY_REQUEST",
	Message:        "주어진 시간동안 너무 많은 요청을 보냈습니다.",
}

var FAILED_DB_PROCESSING = CodeDescription{
	HttpStatusCode: 500,
	Code:           "FAILED_DB_PROCESSING",
	Message:        "잘못된 요청 값으로 처리 중 DB 오류가 발생했습니다.▸ ",
}

var FAILED_INTERNAL_ERROR = CodeDescription{
	HttpStatusCode: 500,
	Code:           "FAILED_INTERNAL_ERROR",
	Message:        "확인되지 않은 오류입니다. ",
}

var FAILED_KEYCLOAK_HANDLING = CodeDescription{
	HttpStatusCode: 500,
	Code:           "FAILED_KEYCLOAK_HANDLING",
	Message:        "인증 시스템 처리 작업이 실패했습니다. 잠시 후 다시 시도해주세요. ",
}
