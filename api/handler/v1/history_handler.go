package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"gorm.io/gorm"
)

func HandleGetPostHistory(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		var items []entity.HistoryItem

		err := db.Model(&entity.DBTablePost{}).
			Select("id, type, title, created_at").
			Where("deleted_at IS NULL").
			Order("created_at DESC").
			Scan(&items).Error
		if err != nil {
			return nil, err
		}

		return entity.NewListResponse(items), nil
	}
}
