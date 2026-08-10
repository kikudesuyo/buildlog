package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleQiitaBatchJob はQiita記事を同期するバッチジョブを実行します。
func HandleQiitaBatchJob(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}

	count, err := service.SyncQiitaArticles(r.Context())
	if err != nil {
		return nil, err
	}

	return entity.NewObjectResponse(map[string]interface{}{
		"synced": count,
	}), nil
}
