package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleSyncQiitaArticleList は管理画面からQiita記事を同期します。
func HandleSyncQiitaArticleList(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}

	count, err := service.SyncQiitaArticleList(r.Context())
	if err != nil {
		return nil, err
	}

	return entity.NewObjectResponse(map[string]interface{}{
		"synced": count,
	}), nil
}
