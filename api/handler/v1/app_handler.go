package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
)

func HandleGetAppList(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	appList, err := service.ListApps(r.Context())
	if err != nil {
		return nil, err
	}
	return entity.NewListResponse(appList), nil
}

func HandleGetApp(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	app, err := service.GetAppByID(r.Context(), id)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(app), nil
}

func HandleCreateApp(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	var req entity.CreateAppRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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
	resp, err := service.UpdateApp(r.Context(), id, req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}

func HandleDeleteApp(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	if err := service.DeleteApp(r.Context(), id); err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(map[string]string{"status": "deleted"}), nil
}
