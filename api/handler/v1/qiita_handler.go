package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleSyncQiitaArticle_List は管理画面からQiita記事を同期します。
func HandleSyncQiitaArticle_List(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}

	count, err := service.SyncQiitaArticle_List(r.Context())
	if err != nil {
		return nil, err
	}

	return entity.NewObjectResponse(map[string]interface{}{
		"synced": count,
	}), nil
}
