package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

func parsePaginationValue(value string) (int, error) {
	if value == "" {
		return 0, nil
	}
	n, err := strconv.Atoi(value)
	if err != nil || n < 0 {
		return 0, http.ErrBodyNotAllowed
	}
	return n, nil
}

// HandleGetTechList はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetTechList(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	query := r.URL.Query()
	if query.Get("all") == "true" {
		if err := handler.ValidateRequestWithAuth(r); err != nil {
			return nil, err
		}
	}
	offset, err := parsePaginationValue(query.Get("offset"))
	if err != nil {
		return nil, err
	}
	limit, err := parsePaginationValue(query.Get("limit"))
	if err != nil {
		return nil, err
	}
	if limit > 0 {
		limit++
	}
	techList, err := service.ListTechs(r.Context(), query.Get("all") == "true", offset, limit, getClientIP(r))
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
	if r.URL.Query().Get("count_view") == "true" {
		if err := service.IncrementTechViews(r.Context(), id); err != nil {
			return nil, err
		}
		tech.Views++
	}
	return entity.NewObjectResponse(tech), nil
}

// HandleCreateTech はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleCreateTech(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
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
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
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
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	if err := service.DeleteTech(r.Context(), id); err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(map[string]string{"status": "deleted"}), nil
}
