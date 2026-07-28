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

func HandleGetTechList() handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		all := r.URL.Query().Get("all") == "true"
		ipAddress := getClientIP(r)
		techs, err := service.ListTechs(r.Context(), all, ipAddress)
		if err != nil {
			return nil, err
		}
		return entity.NewListResponse(techs), nil
	}
}

func HandleGetTech() handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			return nil, err
		}

		ipAddress := getClientIP(r)
		tech, err := service.GetTechByID(r.Context(), id, ipAddress)
		if err != nil {
			return nil, err
		}
		return entity.NewObjectResponse(tech), nil
	}
}

func HandleCreateTech() handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
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
}

func HandleUpdateTech() handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
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
}

func HandleDeleteTech() handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, err
		}

		if err := service.DeleteTech(r.Context(), id); err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(map[string]string{"status": "deleted"}), nil
	}
}
