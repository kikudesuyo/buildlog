package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleGetDiaryList はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetDiaryList(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	all := r.URL.Query().Get("all") == "true"
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if offset < 0 {
		offset = 0
	}
	if limit < 0 {
		limit = 0
	}
	diaryList, err := service.ListDiaries(r.Context(), all, offset, limit, getClientIP(r))
	if err != nil {
		return nil, err
	}
	return entity.NewListResponse(diaryList), nil
}

// HandleGetDiary はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetDiary(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	diary, err := service.GetDiaryByID(r.Context(), id, getClientIP(r))
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(diary), nil
}

// HandleCreateDiary はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleCreateDiary(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	var req entity.CreateDiaryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	if req.Title == "" || req.Content == "" {
		return nil, http.ErrBodyNotAllowed
	}
	resp, err := service.CreateDiary(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}

// HandleUpdateDiary はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleUpdateDiary(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
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
	resp, err := service.UpdateDiary(r.Context(), id, req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(resp), nil
}

// HandleDeleteDiary はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleDeleteDiary(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	if err := service.DeleteDiary(r.Context(), id); err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(map[string]string{"status": "deleted"}), nil
}
