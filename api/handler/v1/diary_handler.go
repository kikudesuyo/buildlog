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

func HandleGetDiaryList(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		diaries, err := service.ListDiaries(r.Context(), db)
		if err != nil {
			return nil, err
		}
		return entity.NewListResponse(diaries), nil
	}
}

func HandleGetDiary(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			return nil, err
		}

		diary, err := service.GetDiaryByID(r.Context(), db, id)
		if err != nil {
			return nil, err
		}
		return entity.NewObjectResponse(diary), nil
	}
}

func HandleCreateDiary(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		var req entity.CreateDiaryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, err
		}

		if req.Title == "" || req.Content == "" {
			return nil, http.ErrBodyNotAllowed
		}

		resp, err := service.CreateDiary(r.Context(), db, req.Title, req.Content)
		if err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(resp), nil
	}
}

func HandleUpdateDiary(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, err
		}

		var req entity.UpdateDiaryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, err
		}

		if req.Title == "" || req.Content == "" {
			return nil, http.ErrBodyNotAllowed
		}

		resp, err := service.UpdateDiary(r.Context(), db, id, req.Title, req.Content)
		if err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(resp), nil
	}
}

func HandleDeleteDiary(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, err
		}

		if err := service.DeleteDiary(r.Context(), db, id); err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(map[string]string{"status": "deleted"}), nil
	}
}
