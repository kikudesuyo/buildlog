package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleGetDeletedPost_List はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleGetDeletedPost_List(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	postList, err := service.GetDeletedPost_List(r.Context())
	if err != nil {
		return nil, err
	}
	return entity.New_ListResponse(postList), nil
}

// HandleRestorePost はHTTPリクエストを受け取り、対応する処理結果を返します。
func HandleRestorePost(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	if err := service.RestorePost(r.Context(), id); err != nil {
		return nil, err
	}
	return entity.NewObjectResponse(map[string]string{"status": "restored"}), nil
}
