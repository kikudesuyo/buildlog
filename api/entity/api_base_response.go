package entity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type BaseResponse struct {
	Meta     ResponseMeta   `json:"meta"`
	Err      *ResponseError `json:"error,omitempty"`
	DataType string         `json:"data_type"`
	Data     any            `json:"data,omitempty"`
	DataList any            `json:"data_list,omitempty"`
}

type ResponseMeta struct {
	StatusCode int  `json:"status_code"`
	Success    bool `json:"success"`
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewObjectResponse(data any) BaseResponse {
	return BaseResponse{Meta: ResponseMeta{StatusCode: http.StatusOK, Success: true}, DataType: "object", Data: data}
}

func NewListResponse(data any) BaseResponse {
	return BaseResponse{Meta: ResponseMeta{StatusCode: http.StatusOK, Success: true}, DataType: "list", DataList: data}
}

func (r BaseResponse) Render(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Meta.StatusCode)
	_, _ = fmt.Fprint(w, r.GetBody())
}

func (r BaseResponse) GetStatusCode() int {
	return r.Meta.StatusCode
}

func (r BaseResponse) GetBody() string {
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(r)
	return buf.String()
}
