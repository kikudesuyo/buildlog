package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

// ProcessFunc is a function to process request and return response data.
// リクエストを処理し、レスポンスデータを返却する関数
type ProcessFunc func(r *http.Request, requestData map[string]interface{}) (http.Handler, error)

// HandleRequestAndResponse handles a request and response.
// リクエストとレスポンスを処理する.
func HandleRequestAndResponse(r *http.Request, w http.ResponseWriter, processFn ProcessFunc) {
	respData, err := handleRequest(r, processFn)
	if err != nil {
		entity.NewErrorResponse(err).ServeHTTP(w, r)
		return
	}

	respData.ServeHTTP(w, r)
}

// ValidateRequestWithAuth はJWT middlewareが保存した認証結果を検証します。
// 公開APIでは呼び出さず、認証が必要なhandler/middlewareから利用します。
func ValidateRequestWithAuth(r *http.Request) error {
	if err, ok := library.CtxGetJWTError(r.Context()); ok && err != nil {
		return err
	}
	return xerror.AuthJWTEmptyToken()
}

// handleRequest handles a request, process it and return response.
// リクエストを受け付け、各処理実行し、レスポンスを返却する.
func handleRequest(r *http.Request, processFn ProcessFunc) (http.Handler, error) {

	// form, body, query stringからパラメータを共通取得
	data, err := getParameters(r)
	if err != nil {
		return nil, err
	}

	// 前処理実行
	r, err = handleRequestBefore(r)
	if err != nil {
		return nil, err
	}

	// 各処理実行
	resp, err := processFn(r, data)

	// 後処理実行
	if afterErr := handleRequestAfter(r, resp); afterErr != nil {
		return nil, afterErr
	}

	return resp, err
}

// getParameters extract all of parameters from Request.
func getParameters(r *http.Request) (map[string]interface{}, error) {
	data := make(map[string]interface{})

	// query string
	params := r.URL.Query()
	for key, vals := range params {
		if len(vals) == 1 {
			data[key] = vals[0]
			continue
		}
		data[key] = vals
	}

	contentType := r.Header.Get("Content-Type")

	// body or form values
	switch {
	case contentType == "application/x-www-form-urlencoded":
		if err := r.ParseForm(); err != nil {
			return nil, err
		}
		for key, val := range r.Form {
			data[key] = val
		}
	case r.Method == "POST",
		r.Method == "PUT",
		contentType == "application/json",
		strings.HasPrefix(contentType, "application/json;"):
		bodyData := make(map[string]interface{})
		buf, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		r.Body = io.NopCloser(bytes.NewReader(buf)) // request.Bodyへストリームデータを書き戻して再利用できるようにする
		err = json.NewDecoder(bytes.NewReader(buf)).Decode(&bodyData)
		if err != nil && err != io.EOF {
			return nil, err
		}
		for key, val := range bodyData {
			data[key] = val
		}
	}

	return data, nil
}
