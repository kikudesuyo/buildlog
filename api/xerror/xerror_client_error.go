package xerror

import (
	"fmt"

	"github.com/morikuni/failure"
)

const (
	CodeClientResourceNotFound        failure.StringCode = "A-0-00001"
	CodeClientInvalidPermission       failure.StringCode = "A-0-00002"
	CodeClientValidation              failure.StringCode = "A-1-10001"
	CodeClientValidationEmail         failure.StringCode = "A-1-10101"
	CodeClientValidationEmailNotMatch failure.StringCode = "A-1-10105"
	CodeClientValidationNoPermission  failure.StringCode = "A-1-10201"
)

// クライアントから指定したデータが存在しない場合のエラー
func ClientResourceNotFoundErr(meta ...map[string]string) error {
	return generateErrFromText("データが存在しません", CodeClientResourceNotFound, meta...)
}

// クライアントが権限を満たさない場合のエラー
func ClientInvalidPermission(meta ...map[string]string) error {
	return generateErrFromText("Invalid permission", CodeClientInvalidPermission, meta...)
}

// クライアント起因によるバリデーションエラー
func ClientValidationErr(err error, meta ...map[string]string) error {
	switch {
	case err == nil:
		return nil
	case IsCustomError(err):
		return AddMetaData(err, meta...)
	}

	msg := fmt.Sprintf("validation error: %s", err.Error())
	return generateErr(err, msg, CodeClientValidation, meta...)
}

// ClientValidationEmail はこの処理に必要な内部処理を実行します。
func ClientValidationEmail(meta ...map[string]string) error {
	return generateErrFromText("メールアドレスの形式が正しくありません", CodeClientValidationEmail, meta...)
}

// ClientValidationEmailDomain はこの処理に必要な内部処理を実行します。
func ClientValidationEmailDomain(domain string, meta ...map[string]string) error {
	return generateErrFromText(fmt.Sprintf("[%s] は利用できるメールアドレスのドメインではありません", domain), CodeClientValidationEmail, meta...)
}

// ClientValidationEmailNotMatch はこの処理に必要な内部処理を実行します。
func ClientValidationEmailNotMatch(email string, meta ...map[string]string) error {
	return generateErrFromText(fmt.Sprintf("登録されているメールアドレスと一致しません: [%s]", email), CodeClientValidationEmailNotMatch, meta...)
}

// ClientValidationNoPermission はこの処理に必要な内部処理を実行します。
func ClientValidationNoPermission(perm string, meta ...map[string]string) error {
	return generateErrFromText(fmt.Sprintf("実行権限がありません: [%s]", perm), CodeClientValidationNoPermission, meta...)
}
