package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
	"gorm.io/gorm"
)

func HandleGetTechList(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		all := r.URL.Query().Get("all") == "true"
		techs, err := service.ListTechs(r.Context(), db, all)
		if err != nil {
			return nil, err
		}
		return entity.NewListResponse(techs), nil
	}
}

func HandleGetTech(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			return nil, err
		}

		tech, err := service.GetTechByID(r.Context(), db, id)
		if err != nil {
			return nil, err
		}
		return entity.NewObjectResponse(tech), nil
	}
}

func HandleCreateTech(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		var req entity.CreateTechRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, err
		}

		if req.Title == "" || req.Content == "" || !entity.IsValidTechCategory(req.Category) {
			return nil, http.ErrBodyNotAllowed
		}

		resp, err := service.CreateTech(r.Context(), db, req)
		if err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(resp), nil
	}
}

func HandleUpdateTech(db *gorm.DB) handler.ProcessFunc {
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

		resp, err := service.UpdateTech(r.Context(), db, id, req)
		if err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(resp), nil
	}
}

func HandleDeleteTech(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, err
		}

		if err := service.DeleteTech(r.Context(), db, id); err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(map[string]string{"status": "deleted"}), nil
	}
}
