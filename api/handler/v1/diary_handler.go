package v1

import (
	"net/http"
	"strconv"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleGetDiaryList はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetDiaryList(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	all := r.URL.Query().Get("all") == "true"
	if all {
		if err := handler.ValidateRequestWithAuth(r); err != nil {
			return nil, err
		}
	}
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	sortBy := r.URL.Query().Get("sort")
	if sortBy != "likes" {
		sortBy = "newest"
	}
	sortOrder := r.URL.Query().Get("order")
	if sortOrder != "asc" {
		sortOrder = "desc"
	}
	if offset < 0 {
		offset = 0
	}
	if limit < 0 {
		limit = 0
	}
	diaryList, err := service.GetDiaryList(r.Context(), all, offset, limit, sortBy, sortOrder, getClientIP(r))
	if err != nil {
		return nil, err
	}
	return entity.NewListResponse(diaryList), nil
}

// HandleGetDiary はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetDiary(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	id, err := parseID(r)
	if err != nil {
		return nil, err
	}
	diary, err := service.GetDiaryByID(r.Context(), id, getClientIP(r))
	if err != nil {
		return nil, err
	}
	if r.URL.Query().Get("count_view") == "true" {
		if err := service.IncrementDiaryViews(r.Context(), id); err != nil {
			return nil, err
		}
		diary.Views++
	}
	return entity.NewObjectResponse(diary), nil
}

// HandleCreateDiary はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleCreateDiary(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	var req entity.CreateDiaryRequest
	if err := decodeJSON(r, &req); err != nil {
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
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	id, err := parseID(r)
	if err != nil {
		return nil, err
	}
	var req entity.UpdateDiaryRequest
	if err := decodeJSON(r, &req); err != nil {
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
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	id, err := parseID(r)
	if err != nil {
		return nil, err
	}
	if err := service.DeleteDiary(r.Context(), id); err != nil {
		return nil, err
	}
	return deletedResponse("deleted"), nil
}
