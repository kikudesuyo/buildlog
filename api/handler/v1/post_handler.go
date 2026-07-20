package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

func HandleGetPostList(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
	return entity.NewListResponse(service.GetPostList(r.Context())), nil
}
