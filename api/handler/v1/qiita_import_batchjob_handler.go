package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/external"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/service"
)

// HandleQiitaImportBatchJob はQiita記事を同期するバッチジョブを実行します。
func HandleQiitaImportBatchJob(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
	if err := handler.ValidateRequestWithAuth(r); err != nil {
		return nil, err
	}

	qiitaClient := external.NewQiitaClient(nil)
	count, err := service.ImportQiitaItems(r.Context(), library.GetDB(r.Context()), qiitaClient, qiitaClient)
	if err != nil {
		return nil, err
	}

	return entity.NewObjectResponse(map[string]interface{}{
		"imported": count,
	}), nil
}
