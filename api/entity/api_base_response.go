package entity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kikudesuyo/buildlog/api/xerror"
)

type BaseResponse struct {
	Meta     ResponseMeta   `json:"meta"`
	Err      *ResponseError `json:"error,omitempty"`
	DataType string         `json:"data_type"`
	Data     any            `json:"data,omitempty"`
	DataList any            `json:"data_list,omitempty"`
	HasMore  *bool          `json:"has_more,omitempty"`
}

type ResponseMeta struct {
	StatusCode int  `json:"status_code"`
	Success    bool `json:"success"`
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewObjectResponse は値を生成します。
func NewObjectResponse(data any) BaseResponse {
	return BaseResponse{Meta: ResponseMeta{StatusCode: http.StatusOK, Success: true}, DataType: "object", Data: data}
}

// NewListResponse は値を生成します。
func NewListResponse(data any) BaseResponse {
	return BaseResponse{Meta: ResponseMeta{StatusCode: http.StatusOK, Success: true}, DataType: "list", DataList: data}
}

// NewListResponseWithHasMore はページング情報付きの一覧レスポンスを生成します。
func NewListResponseWithHasMore(data any, hasMore bool) BaseResponse {
	return BaseResponse{Meta: ResponseMeta{StatusCode: http.StatusOK, Success: true}, DataType: "list", DataList: data, HasMore: &hasMore}
}

// NewErrorResponse は値を生成します。
func NewErrorResponse(err error) BaseResponse {
	return BaseResponse{
		Meta: ResponseMeta{
			StatusCode: xerror.GetHTTPErrorCode(err),
		},
		Err: &ResponseError{
			Code:    xerror.GetStringCode(err),
			Message: xerror.GetErrorMessage(err),
		},
		DataType: "error",
	}
}

// Render はこの処理に必要な内部処理を実行します。
func (r BaseResponse) Render(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Meta.StatusCode)
	_, _ = fmt.Fprint(w, r.GetBody())
}

// ServeHTTP はこの処理に必要な内部処理を実行します。
func (r BaseResponse) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	r.Render(w)
}

// GetStatusCode はデータを取得します。
func (r BaseResponse) GetStatusCode() int {
	return r.Meta.StatusCode
}

// GetBody はデータを取得します。
func (r BaseResponse) GetBody() string {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(r)
	return buf.String()
}
