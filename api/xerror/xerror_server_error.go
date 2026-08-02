package xerror

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/morikuni/failure"
)

const (
	// E-0-0~1... は基盤部分のエラーコード
	CodeAppLib           failure.StringCode = "E-0-00003"
	CodeAppDebug         failure.StringCode = "E-0-00004"
	CodeInvalidSetting   failure.StringCode = "E-0-00005"
	CodeAppPanic         failure.StringCode = "E-0-00099"
	CodeDBInRepository   failure.StringCode = "E-0-10101"
	CodeDB               failure.StringCode = "E-0-10001"
	CodeDBInit           failure.StringCode = "E-0-10002"
	CodeDBPrimaryKey     failure.StringCode = "E-0-10003"
	CodeDBCommit         failure.StringCode = "E-0-10004"
	CodeDBDuplicateEntry failure.StringCode = "E-0-10005"
	CodeDBSchemaColumn   failure.StringCode = "E-0-10006"

	// E-0-2 以降は外部API・システム関連のエラーコード
	CodeGCP           failure.StringCode = "E-0-20001"
	CodeStorage       failure.StringCode = "E-0-20101"
	CodeStorageUpload failure.StringCode = "E-0-20102"
	CodeStorageDelete failure.StringCode = "E-0-20103"
	CodeStorageCopy   failure.StringCode = "E-0-20104"
	CodePinpoint      failure.StringCode = "E-0-40001"
	CodeEmailTemplate failure.StringCode = "E-0-40002"
	CodeEmailSend     failure.StringCode = "E-0-40003"
	CodeCache         failure.StringCode = "E-0-50001"
	CodeCacheGet      failure.StringCode = "E-0-50002"

	CodeTodo            failure.StringCode = "E-1-19999"
	CodeDataEmpty       failure.StringCode = "E-1-10003"
	CodeTargetDataEmpty failure.StringCode = "E-1-10004"
	CodeImageDataEmpty  failure.StringCode = "E-1-10005"

	CodeJSONUnmarshalStruct failure.StringCode = "E-2-10001"
	CodeTemplateParse       failure.StringCode = "E-2-20201"
	CodeTemplateExecute     failure.StringCode = "E-2-20202"

	// LLM関連のエラーコード
	CodeLLMClientInit      failure.StringCode = "E-0-30001"
	CodeLLMGenerateContent failure.StringCode = "E-0-30002"
	CodeLLMEmptyResponse   failure.StringCode = "E-0-30003"

	// Google Sheets API関連のエラーコード
	CodeSheetsServiceInit   failure.StringCode = "E-0-21001"
	CodeSheetsRead          failure.StringCode = "E-0-21002"
	CodeSheetsReadRateLimit failure.StringCode = "E-0-21003"
	CodeSheetsWrite         failure.StringCode = "E-0-21004"
	CodeSheetsInvalidEnv    failure.StringCode = "E-0-21005"

	// CSV処理関連のエラーコード
	CodeCSVParse             failure.StringCode = "E-1-20001"
	CodeCSVInvalidFieldCount failure.StringCode = "E-1-20002"
	CodeCSVDataConversion    failure.StringCode = "E-1-20003"
	CodeCSVWrite             failure.StringCode = "E-1-20004"

	// PDF処理関連のエラーコード
	CodePDFDecode   failure.StringCode = "E-1-30001"
	CodePDFAnalysis failure.StringCode = "E-1-30002"
	CodePDFHash     failure.StringCode = "E-1-30003"

	// テキスト処理関連のエラーコード
	CodeTextAnalysis failure.StringCode = "E-1-40001"

	// キャッシュ関連のエラーコード
	CodeCacheSet         failure.StringCode = "E-0-50003"
	CodeCacheInvalidType failure.StringCode = "E-0-50004"

	// BigQuery関連のエラーコード
	CodeBigQuery     failure.StringCode = "E-0-22001"
	CodeBigQueryLoad failure.StringCode = "E-0-22002"
)

var reJSONStructField = regexp.MustCompile(`struct field .*\.([a-zA-Z_]+) of type`)

// Todo は未実装のエラーを表すエラー
func Todo(meta ...map[string]string) error {
	return generateErrFromText("todo error", CodeTodo, meta...)
}

// TodoErr は未実装のエラーを表すエラー
func TodoErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "todo error", CodeTodo, meta...)
}

// DataEmpty はデータが空のエラーを表すエラー
func DataEmpty(meta ...map[string]string) error {
	return generateErrFromText("data empty error", CodeDataEmpty, meta...)
}

// UnknownLibraryErr は未知のライブラリエラーを表すエラー
func UnknownLibraryErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "unknown app component error", CodeAppLib, meta...)
}

// UnknownPanicError は対応するエラーを生成します。
func UnknownPanicError(meta ...map[string]string) error {
	return generateErrFromText("unknown system error", CodeAppPanic, meta...)
}

// AppDebug はこの処理に必要な内部処理を実行します。
func AppDebug(meta ...map[string]string) error {
	return generateErrFromText("technical error", CodeAppDebug, meta...)
}

// InvalidSetting は設定エラーを表すエラー
func InvalidSetting(meta ...map[string]string) error {
	return generateErrFromText("invalid setting error", CodeInvalidSetting, meta...)
}

// JSONUnmarshalStructErr はJSONのUnmarshalエラーを表すエラー
func JSONUnmarshalStructErr(err error, meta ...map[string]string) error {
	if err == nil {
		return nil
	}

	results := reJSONStructField.FindAllStringSubmatch(err.Error(), -1)
	if len(results) == 0 {
		return UnknownLibraryErr(err)
	}
	res := results[0]
	if len(res) != 2 {
		return UnknownLibraryErr(err)
	}

	return generateErrFromText(fmt.Sprintf("%s の型が不正です.", res[1]), CodeJSONUnmarshalStruct, meta...)
}

// レポジトリ層でのエラー用（stack trace判定のため）
func UnknownDBInRepositoryErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErrWithStackDepth(err, "unknown database error", CodeDBInRepository, 1, meta...)
}

// UnknownDBErr は対応するエラーを生成します。
func UnknownDBErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "unknown database error", CodeDB, meta...)
}

// DBInitErr は対応するエラーを生成します。
func DBInitErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "unknown database init error", CodeDBInit, meta...)
}

// DBPrimaryKey はこの処理に必要な内部処理を実行します。
func DBPrimaryKey(meta ...map[string]string) error {
	return generateErr(errors.New("primary key must be null on creation"), "unknown database error", CodeDBPrimaryKey, meta...)
}

// DBCommitErr は対応するエラーを生成します。
func DBCommitErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "unknown database commit error", CodeDBCommit, meta...)
}

// DBDuplicateEntryErr は対応するエラーを生成します。
func DBDuplicateEntryErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "the data already exists", CodeDBDuplicateEntry, meta...)
}

// DBSchemaColumnErr は対応するエラーを生成します。
func DBSchemaColumnErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "database schema column error (e.g. default value)", CodeDBSchemaColumn, meta...)
}

// UnknownPinpointErr は対応するエラーを生成します。
func UnknownPinpointErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "unknown email error", CodePinpoint, meta...)
}

// UnknownEmailTemplateErr は対応するエラーを生成します。
func UnknownEmailTemplateErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "unknown email error", CodeEmailTemplate, meta...)
}

// UnknownEmailSendErr は対応するエラーを生成します。
func UnknownEmailSendErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "unknown email error", CodeEmailSend, meta...)
}

// CacheGetErr は対応するエラーを生成します。
func CacheGetErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "cache get error", CodeCacheGet, meta...)
}

// CacheSetErr は対応するエラーを生成します。
func CacheSetErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "cache set error", CodeCacheSet, meta...)
}

// CacheInvalidTypeErr は対応するエラーを生成します。
func CacheInvalidTypeErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "cache invalid type error", CodeCacheInvalidType, meta...)
}

// LLMClientInitErr はLLMクライアントの初期化エラーを表すエラー
func LLMClientInitErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "LLM client initialization error", CodeLLMClientInit, meta...)
}

// LLMGenerateContentErr はLLMのコンテンツ生成エラーを表すエラー
func LLMGenerateContentErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "LLM content generation error", CodeLLMGenerateContent, meta...)
}

// LLMEmptyResponse はLLMの空レスポンスエラーを表すエラー
func LLMEmptyResponse(meta ...map[string]string) error {
	return generateErrFromText("LLM empty response error", CodeLLMEmptyResponse, meta...)
}

// SheetsServiceInitErr はGoogle Sheetsサービスの初期化エラーを表すエラー
func SheetsServiceInitErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "Sheets service initialization error", CodeSheetsServiceInit, meta...)
}

// SheetsReadErr はGoogle Sheetsの読み込みエラーを表すエラー
func SheetsReadErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "Sheets read error", CodeSheetsRead, meta...)
}

// SheetsReadRateLimitErr はGoogle Sheetsの読み込みレート制限エラーを表すエラー
func SheetsReadRateLimitErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "Sheets read rate limit error", CodeSheetsReadRateLimit, meta...)
}

// SheetsWriteErr はGoogle Sheetsの書き込みエラーを表すエラー
func SheetsWriteErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "Sheets write error", CodeSheetsWrite, meta...)
}

// SheetsInvalidEnv はGoogle Sheetsの環境変数エラーを表すエラー
func SheetsInvalidEnv(meta ...map[string]string) error {
	return generateErrFromText("Sheets environment variable error", CodeSheetsInvalidEnv, meta...)
}

// CSVParseErr はCSVのパースエラーを表すエラー
func CSVParseErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "CSV parse error", CodeCSVParse, meta...)
}

// CSVInvalidFieldCount はCSVのフィールド数エラーを表すエラー
func CSVInvalidFieldCount(meta ...map[string]string) error {
	return generateErrFromText("CSV invalid field count error", CodeCSVInvalidFieldCount, meta...)
}

// CSVDataConversionErr はCSVのデータ変換エラーを表すエラー
func CSVDataConversionErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "CSV data conversion error", CodeCSVDataConversion, meta...)
}

// CSVWriteErr はCSVの書き込みエラーを表すエラー
func CSVWriteErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "CSV write error", CodeCSVWrite, meta...)
}

// PDFDecodeErr はPDFのデコードエラーを表すエラー
func PDFDecodeErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "PDF decode error", CodePDFDecode, meta...)
}

// PDFAnalysisErr はPDFの解析エラーを表すエラー
func PDFAnalysisErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "PDF analysis error", CodePDFAnalysis, meta...)
}

// TextAnalysisErr はTextの解析エラーを表すエラー
func TextAnalysisErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "Text analysis error", CodeTextAnalysis, meta...)
}

// PDFHashErr はPDFのハッシュ生成エラーを表すエラー
func PDFHashErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "PDF hash generation error", CodePDFHash, meta...)
}

// StorageUploadErr はGCSアップロードエラーを表すエラー
func StorageUploadErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "storage upload error", CodeStorageUpload, meta...)
}

// BigQueryErr はBigQueryの一般エラーを表すエラー
func BigQueryErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "BigQuery error", CodeBigQuery, meta...)
}

// BigQueryLoadErr はBigQueryのデータロードエラーを表すエラー
func BigQueryLoadErr(err error, meta ...map[string]string) error {
	if IsCustomError(err) {
		return AddMetaData(err, meta...)
	}
	return generateErr(err, "BigQuery load error", CodeBigQueryLoad, meta...)
}
