package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleGetPostHistory はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetPostHistory(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	itemList, err := service.GetPostHistory_List(r.Context())
	if err != nil {
		return nil, err
	}
	return entity.New_ListResponse(itemList), nil
}
