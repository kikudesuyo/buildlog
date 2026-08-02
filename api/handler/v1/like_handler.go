package v1

import (
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
)

// getClientIP はこの処理に必要な内部処理を実行します。
func getClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

// HandlePostLike はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandlePostLike(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	postID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	resp, err := service.LikePost(r.Context(), postID, getClientIP(r))
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}

// HandleDeleteLike はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleDeleteLike(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	postID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	resp, err := service.UnlikePost(r.Context(), postID, getClientIP(r))
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}

// HandleGetLikeStatus はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetLikeStatus(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	postID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	resp, err := service.GetLikeStatus(r.Context(), postID, getClientIP(r))
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}
