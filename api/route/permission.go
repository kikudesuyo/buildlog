package route

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/handler"
)

// handleFunc は権限チェック不要の通常ハンドラを HTTPハンドラに変換する.
func handleFunc(processFn handler.ProcessFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.HandleRequestAndResponse(r, w, processFn)
	}
}
