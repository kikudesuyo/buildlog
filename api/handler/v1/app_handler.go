package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleGetAppList はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetAppList(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	appList, err := service.ListApps(r.Context(), service.DatabaseFromContext(r.Context()))
	if err != nil {
		return nil, err
	}
	return entity.NewListResponse(appList), nil
}

// HandleGetApp はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetApp(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	app, err := service.GetAppByID(r.Context(), service.DatabaseFromContext(r.Context()), id)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(app), nil
}

// HandleCreateApp はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleCreateApp(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	var req entity.CreateAppRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	if req.Slug == "" || req.Name == "" {
		return nil, http.ErrBodyNotAllowed
	}
	resp, err := service.CreateApp(r.Context(), service.DatabaseFromContext(r.Context()), req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}

// HandleUpdateApp はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleUpdateApp(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	var req entity.UpdateAppRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	if req.Slug == "" || req.Name == "" {
		return nil, http.ErrBodyNotAllowed
	}
	resp, err := service.UpdateApp(r.Context(), service.DatabaseFromContext(r.Context()), id, req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}

// HandleDeleteApp はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleDeleteApp(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	if err := service.DeleteApp(r.Context(), service.DatabaseFromContext(r.Context()), id); err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(map[string]string{"status": "deleted"}), nil
}
