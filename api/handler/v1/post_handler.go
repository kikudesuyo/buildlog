package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"github.com/kikudesuyo/buildlog/api/service"
	"gorm.io/gorm"
)

func HandleGetPostList(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		posts, err := service.GetPostList(r.Context(), db)
		if err != nil {
			return nil, err
		}
		return entity.NewListResponse(posts), nil
	}
}
