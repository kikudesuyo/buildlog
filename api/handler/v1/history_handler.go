package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleGetPostHistory はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetPostHistory(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	itemList, err := service.ListPostHistory(r.Context())
	if err != nil {
		return nil, err
	}
	return entity.NewListResponse(itemList), nil
}
