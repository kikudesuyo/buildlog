package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleGetTechList はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetTechList(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	techList, err := service.ListTechs(r.Context(), r.URL.Query().Get("all") == "true", r.URL.Query().Get("q"), getClientIP(r))
	if err != nil {
		return nil, err
	}
	return entity.NewListResponse(techList), nil
}

// HandleGetTech はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetTech(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	tech, err := service.GetTechByID(r.Context(), id, getClientIP(r))
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(tech), nil
}

// HandleCreateTech はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleCreateTech(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	var req entity.CreateTechRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	if req.Title == "" || req.Content == "" || !entity.IsValidTechCategory(req.Category) {
		return nil, http.ErrBodyNotAllowed
	}
	resp, err := service.CreateTech(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}

// HandleUpdateTech はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleUpdateTech(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	var req entity.UpdateTechRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	if req.Title == "" || req.Content == "" || !entity.IsValidTechCategory(req.Category) {
		return nil, http.ErrBodyNotAllowed
	}
	resp, err := service.UpdateTech(r.Context(), id, req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}

// HandleDeleteTech はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleDeleteTech(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	if err := service.DeleteTech(r.Context(), id); err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(map[string]string{"status": "deleted"}), nil
}
