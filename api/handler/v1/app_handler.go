package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleGetAppList はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetAppList(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	appList, err := service.GetAppList(r.Context())
	if err != nil {
		return nil, err
	}
	return entity.NewListResponse(appList), nil
}

// HandleGetApp はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetApp(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	id, err := parseID(r)
	if err != nil {
		return nil, err
	}
	app, err := service.GetAppByID(r.Context(), id)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(app), nil
}

// HandleCreateApp はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleCreateApp(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	var req entity.CreateAppRequest
	if err := decodeJSON(r, &req); err != nil {
		return nil, err
	}
	if req.Slug == "" || req.Name == "" {
		return nil, http.ErrBodyNotAllowed
	}
	resp, err := service.CreateApp(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}

// HandleUpdateApp はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleUpdateApp(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	id, err := parseID(r)
	if err != nil {
		return nil, err
	}
	var req entity.UpdateAppRequest
	if err := decodeJSON(r, &req); err != nil {
		return nil, err
	}
	if req.Slug == "" || req.Name == "" {
		return nil, http.ErrBodyNotAllowed
	}
	resp, err := service.UpdateApp(r.Context(), id, req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}

// HandleDeleteApp はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleDeleteApp(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	id, err := parseID(r)
	if err != nil {
		return nil, err
	}
	if err := service.DeleteApp(r.Context(), id); err != nil {
		return nil, err
	}
	return deletedResponse("deleted"), nil
}
