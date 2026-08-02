package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
)

func HandleGetPostHistory(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	items, err := service.ListPostHistory(r.Context())
	if err != nil {
		return nil, err
	}
	return entity.NewListResponse(items), nil
}
