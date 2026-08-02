package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
)

func HandleGetAnalytics(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	analytics, err := service.GetAnalytics(r.Context())
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(analytics), nil
}
