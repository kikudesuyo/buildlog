package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/kikudesuyo/buildlog/api/entity"
)

// ProcessFunc is a function to process request and return response data.
// リクエストを処理し、レスポンスデータを返却する関数
type ProcessFunc func(r *http.Request, requestData map[string]interface{}) (Renderer, error)

// HandleRequestAndResponse handles a request and response.
// リクエストとレスポンスを処理する.
func HandleRequestAndResponse(r *http.Request, w http.ResponseWriter, processFn ProcessFunc) {
	respData, err := handleRequest(r, processFn)
	if err != nil {
		entity.NewErrorResponse(err).Render(w)
		return
	}

	respData.Render(w)
}

// handleRequest handles a request, process it and return response.
// リクエストを受け付け、各処理実行し、レスポンスを返却する.
func handleRequest(r *http.Request, processFn ProcessFunc) (Renderer, error) {

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
	if err := handleRequestAfter(r, resp); err != nil {
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
		buf, _ := io.ReadAll(r.Body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))
		r.Body = rdr2 // request.Bodyへストリームデータを書き戻して再利用できるようにする
		err := json.NewDecoder(rdr1).Decode(&bodyData)
		if err != nil && err != io.EOF {
			return nil, err
		}
		for key, val := range bodyData {
			data[key] = val
		}
	}

	return data, nil
}
