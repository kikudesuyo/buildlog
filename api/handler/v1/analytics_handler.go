package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleGetAnalytics はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetAnalytics(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	analytics, err := service.GetAnalytics(r.Context(), service.DatabaseFromContext(r.Context()))
	if err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(analytics), nil
}
