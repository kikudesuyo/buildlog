package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"gorm.io/gorm"
)

func HandleGetDeletedPosts(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		var postList []entity.DBTablePost
		err := db.WithContext(r.Context()).
			Unscoped().
			Where("deleted_at IS NOT NULL").
			Order("deleted_at DESC").
			Find(&postList).Error
		if err != nil {
			return nil, err
		}

		return entity.NewListResponse(postList), nil
	}
}

func HandleRestorePost(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			return nil, err
		}

		err = db.WithContext(r.Context()).
			Unscoped().
			Model(&entity.DBTablePost{}).
			Where("id = ?", id).
			Update("deleted_at", nil).Error
		if err != nil {
			return nil, err
		}

		return entity.NewObjectResponse(map[string]string{"status": "restored"}), nil
	}
}
