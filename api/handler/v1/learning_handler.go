package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

func HandleFetchCurrentLearnings(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	periodType, _ := requestData["period_type"].(string)
	if periodType == "" {
		periodType = service.DailyLearning
	}
	items, err := service.FetchCurrentLearnings(r.Context(), periodType, time.Now())
	if err != nil {
		return nil, err
	}
	return entity.NewListResponse(items), nil
}

func HandleCreateDailyLearning(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	var req entity.CreateLearningRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, xerror.ClientValidationErr(err)
	}
	item, err := service.CreateDailyLearning(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(item), nil
}

func HandleCreateLearningSummary(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	periodType := chi.URLParam(r, "period_type")
	var req entity.CreateLearningSummaryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, xerror.ClientValidationErr(err)
	}
	item, err := service.CreateLearningSummary(r.Context(), periodType, req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(item), nil
}
