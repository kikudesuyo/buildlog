package v1

import (
	"encoding/json"
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

// HandleGetCurrentGoals はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetCurrentGoals(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	goals, err := service.GetCurrentGoals(r.Context())
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(goals), nil
}

// HandleGetGoalHistory はHTTPリクエストを受け取り、目標履歴を返します。
func HandleGetGoalHistory(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	history, err := service.GetGoalHistory(r.Context())
	if err != nil {
		return nil, err
	}
	return entity.NewListResponse(history), nil
}

// HandleSaveCurrentGoals はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleSaveCurrentGoals(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	var req entity.SaveGoalsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, xerror.ClientValidationErr(err)
	}
	goals, err := service.SaveCurrentGoals(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(goals), nil
}
