package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

func HandleListDiaryEntries(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
	return entity.NewListResponse(service.ListDiaryEntries(r.Context())), nil
}

func HandleGetTechFeed(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
	return entity.NewObjectResponse(service.GetTechFeed(r.Context())), nil
}

func HandleGetProfileData(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
	return entity.NewObjectResponse(service.GetProfileData(r.Context())), nil
}
