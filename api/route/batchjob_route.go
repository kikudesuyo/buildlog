package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	v1 "github.com/kikudesuyo/buildlog/api/handler/v1"
)

// NewBatchJobRouter はバッチジョブ用のルーターを生成します。
func NewBatchJobRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/qiita/import", handleFunc(v1.HandleQiitaImportBatchJob))
	return r
}
