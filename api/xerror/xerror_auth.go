package xerror

import (
	"errors"

	"github.com/morikuni/failure"
)

const (
	CodeAuthError                          failure.StringCode = "C-0-00001"
	CodeAuthMissingHeaderError             failure.StringCode = "C-0-10001"
	CodeAuthInvalidHeaderError             failure.StringCode = "C-0-10002"
	CodeAuthJWTEmptyTokenError             failure.StringCode = "C-0-20001"
	CodeAuthJWTInvalidTokenError           failure.StringCode = "C-0-20002"
	CodeAuthJWTInvalidClaimsBodyError      failure.StringCode = "C-0-20003"
	CodeAuthJWTInvalidClaimsTypeError      failure.StringCode = "C-0-20004"
	CodeAuthJWTInvalidClaimsDataError      failure.StringCode = "C-0-20005"
	CodeAuthJWTInvalidClaimsIssuerError    failure.StringCode = "C-0-20006"
	CodeAuthJWTInvalidClaimsIssuedAtError  failure.StringCode = "C-0-20007"
	CodeAuthJWTInvalidClaimsNotBeforeError failure.StringCode = "C-0-20008"
	CodeAuthJWTExpiredTokenError           failure.StringCode = "C-0-30001"
	CodeAuthEmptyAccountError              failure.StringCode = "C-0-40001"
	CodeAuthDisabledAccountError           failure.StringCode = "C-0-40002"
	CodeAuthTooManyTryError                failure.StringCode = "C-0-41001"
	CodeAuthAdminSessionInvalidError       failure.StringCode = "C-0-50001"

	CodeAuthDummyError failure.StringCode = "C-1-00001"
)

// AuthGeneralErr は対応するエラーを生成します。
func AuthGeneralErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "auth error", CodeAuthError, meta...)
}

// AuthMissingHeader はこの処理に必要な内部処理を実行します。
func AuthMissingHeader(meta ...map[string]string) error {
	return generateErrFromText("Cannot find 'Authorization' header", CodeAuthMissingHeaderError, meta...)
}

// AuthInvalidHeader はこの処理に必要な内部処理を実行します。
func AuthInvalidHeader(meta ...map[string]string) error {
	return generateErrFromText("Invalid data on 'Authorization' header", CodeAuthInvalidHeaderError, meta...)
}

// AuthJWTEmptyToken はこの処理に必要な内部処理を実行します。
func AuthJWTEmptyToken(meta ...map[string]string) error {
	return generateErrFromText("Empty auth token", CodeAuthJWTEmptyTokenError, meta...)
}

// AuthJWTExpiredToken はこの処理に必要な内部処理を実行します。
func AuthJWTExpiredToken(meta ...map[string]string) error {
	return generateErrFromText("Expired auth token", CodeAuthJWTExpiredTokenError, meta...)
}

// AuthJWTInvalidTokenErr は対応するエラーを生成します。
func AuthJWTInvalidTokenErr(err error, meta ...map[string]string) error {
	return generateErr(err, "Invalid auth token", CodeAuthJWTInvalidTokenError, meta...)
}

// AuthJWTInvalidClaimsBody はこの処理に必要な内部処理を実行します。
func AuthJWTInvalidClaimsBody(meta ...map[string]string) error {
	return generateErrFromText("Invalid claims body", CodeAuthJWTInvalidClaimsBodyError, meta...)
}

// AuthJWTInvalidClaimsType はこの処理に必要な内部処理を実行します。
func AuthJWTInvalidClaimsType(meta ...map[string]string) error {
	return generateErrFromText("Invalid claims type", CodeAuthJWTInvalidClaimsTypeError, meta...)
}

// AuthJWTInvalidClaimsData はこの処理に必要な内部処理を実行します。
func AuthJWTInvalidClaimsData(err error, meta ...map[string]string) error {
	return generateErr(err, "Invalid claims data", CodeAuthJWTInvalidClaimsDataError, meta...)
}

// AuthJWTInvalidClaimsIssuer はこの処理に必要な内部処理を実行します。
func AuthJWTInvalidClaimsIssuer(meta ...map[string]string) error {
	return generateErrFromText("Invalid claims content", CodeAuthJWTInvalidClaimsIssuerError, meta...)
}

// AuthJWTInvalidClaimsIssuedAt はこの処理に必要な内部処理を実行します。
func AuthJWTInvalidClaimsIssuedAt(meta ...map[string]string) error {
	return generateErrFromText("Invalid claims", CodeAuthJWTInvalidClaimsIssuerError, meta...)
}

// AuthJWTInvalidClaimsNotBefore はこの処理に必要な内部処理を実行します。
func AuthJWTInvalidClaimsNotBefore(meta ...map[string]string) error {
	return generateErrFromText("Invalid claims", CodeAuthJWTInvalidClaimsNotBeforeError, meta...)
}

// AuthJWTInvalidClaimsExpired はこの処理に必要な内部処理を実行します。
func AuthJWTInvalidClaimsExpired(meta ...map[string]string) error {
	return generateErr(errors.New("expired token"), "auth error", CodeAuthError, meta...) // ECODE=general
}

// AuthJWTInvalidClaimsNoOperatorID はこの処理に必要な内部処理を実行します。
func AuthJWTInvalidClaimsNoOperatorID(meta ...map[string]string) error {
	return generateErr(errors.New("OperatorID does not exists in JWT"), "Invalid claims", CodeAuthError, meta...) // ECODE=general
}

// AuthJWTInvalidClaimsInvalidOperatorID はこの処理に必要な内部処理を実行します。
func AuthJWTInvalidClaimsInvalidOperatorID(err error, meta ...map[string]string) error {
	return generateErr(errors.New("invalid OperatorID in JWT: ["+err.Error()+"]"), "Invalid claims", CodeAuthError, meta...) // ECODE=general
}

// AuthEmptyAccount はこの処理に必要な内部処理を実行します。
func AuthEmptyAccount(meta ...map[string]string) error {
	return generateErr(errors.New("empty account"), "auth error", CodeAuthEmptyAccountError, meta...)
}

// AuthDisabledAccount はこの処理に必要な内部処理を実行します。
func AuthDisabledAccount(meta ...map[string]string) error {
	return generateErr(errors.New("disabled account"), "auth error", CodeAuthDisabledAccountError, meta...)
}

// AuthTooManyTry はこの処理に必要な内部処理を実行します。
func AuthTooManyTry(meta ...map[string]string) error {
	return generateErrFromText("認証の上限を超えました。しばらくしてからお試しください。", CodeAuthTooManyTryError, meta...)
}

func AuthAdminSessionInvalid(meta ...map[string]string) error {
	return generateErrFromText("Invalid admin session", CodeAuthAdminSessionInvalidError, meta...)
}

// AuthDummyErr は対応するエラーを生成します。
func AuthDummyErr(msg string, meta ...map[string]string) error {
	return generateErr(errors.New(msg), "Invalid data", CodeAuthDummyError, meta...)
}
