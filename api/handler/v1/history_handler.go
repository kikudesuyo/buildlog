package v1

import (
	"net/http"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"gorm.io/gorm"
)

func HandleGetPostHistory(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (http.Handler, error) {
		itemList := make([]entity.HistoryItem, 0)

		err := db.Model(&entity.DBTablePost{}).
			Select("id, type, title, created_at").
			Where("deleted_at IS NULL").
			Order("created_at DESC").
			Scan(&itemList).Error
		if err != nil {
			return nil, err
		}

		return entity.NewListResponse(itemList), nil
	}
}
