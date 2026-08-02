package v1

import (
	"encoding/json"
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

func HandleGetCurrentGoals(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	goals, err := service.GetCurrentGoals(r.Context())
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(goals), nil
}

func HandleSaveCurrentGoals(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
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
