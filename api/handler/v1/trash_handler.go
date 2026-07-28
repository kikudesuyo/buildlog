package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
)

func HandleGetDeletedPosts() handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		posts, err := service.ListDeletedPosts(r.Context())
		if err != nil {
			return nil, err
		}

		return entity.NewListResponse(posts), nil
	}
}

func HandleRestorePost() handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			return nil, err
		}

		err = service.RestorePost(r.Context(), id)
		if err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(map[string]string{"status": "restored"}), nil
	}
}
