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

	CodeAuthDummyError failure.StringCode = "C-1-00001"
)

func AuthGeneralErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "auth error", CodeAuthError, meta...)
}

func AuthMissingHeader(meta ...map[string]string) error {
	return generateErrFromText("Cannot find 'Authorization' header", CodeAuthMissingHeaderError, meta...)
}

func AuthInvalidHeader(meta ...map[string]string) error {
	return generateErrFromText("Invalid data on 'Authorization' header", CodeAuthInvalidHeaderError, meta...)
}

func AuthJWTEmptyToken(meta ...map[string]string) error {
	return generateErrFromText("Empty auth token", CodeAuthJWTEmptyTokenError, meta...)
}

func AuthJWTExpiredToken(meta ...map[string]string) error {
	return generateErrFromText("Expired auth token", CodeAuthJWTExpiredTokenError, meta...)
}

func AuthJWTInvalidTokenErr(err error, meta ...map[string]string) error {
	return generateErr(err, "Invalid auth token", CodeAuthJWTInvalidTokenError, meta...)
}

func AuthJWTInvalidClaimsBody(meta ...map[string]string) error {
	return generateErrFromText("Invalid claims body", CodeAuthJWTInvalidClaimsBodyError, meta...)
}

func AuthJWTInvalidClaimsType(meta ...map[string]string) error {
	return generateErrFromText("Invalid claims type", CodeAuthJWTInvalidClaimsTypeError, meta...)
}

func AuthJWTInvalidClaimsData(err error, meta ...map[string]string) error {
	return generateErr(err, "Invalid claims data", CodeAuthJWTInvalidClaimsDataError, meta...)
}

func AuthJWTInvalidClaimsIssuer(meta ...map[string]string) error {
	return generateErrFromText("Invalid claims content", CodeAuthJWTInvalidClaimsIssuerError, meta...)
}

func AuthJWTInvalidClaimsIssuedAt(meta ...map[string]string) error {
	return generateErrFromText("Invalid claims", CodeAuthJWTInvalidClaimsIssuerError, meta...)
}

func AuthJWTInvalidClaimsNotBefore(meta ...map[string]string) error {
	return generateErrFromText("Invalid claims", CodeAuthJWTInvalidClaimsNotBeforeError, meta...)
}

func AuthJWTInvalidClaimsExpired(meta ...map[string]string) error {
	return generateErr(errors.New("expired token"), "auth error", CodeAuthError, meta...) // ECODE=general
}

func AuthJWTInvalidClaimsNoOperatorID(meta ...map[string]string) error {
	return generateErr(errors.New("OperatorID does not exists in JWT"), "Invalid claims", CodeAuthError, meta...) // ECODE=general
}

func AuthJWTInvalidClaimsInvalidOperatorID(err error, meta ...map[string]string) error {
	return generateErr(errors.New("invalid OperatorID in JWT: ["+err.Error()+"]"), "Invalid claims", CodeAuthError, meta...) // ECODE=general
}

func AuthEmptyAccount(meta ...map[string]string) error {
	return generateErr(errors.New("empty account"), "auth error", CodeAuthEmptyAccountError, meta...)
}

func AuthDisabledAccount(meta ...map[string]string) error {
	return generateErr(errors.New("disabled account"), "auth error", CodeAuthDisabledAccountError, meta...)
}

func AuthTooManyTry(meta ...map[string]string) error {
	return generateErrFromText("認証の上限を超えました。しばらくしてからお試しください。", CodeAuthTooManyTryError, meta...)
}

func AuthDummyErr(msg string, meta ...map[string]string) error {
	return generateErr(errors.New(msg), "Invalid data", CodeAuthDummyError, meta...)
}
