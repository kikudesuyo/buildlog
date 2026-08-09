package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// ListPostHistory は削除されていない投稿履歴を取得します。
func ListPostHistory(ctx context.Context, db *gorm.DB) ([]entity.HistoryItem, error) {
	itemList := make([]entity.HistoryItem, 0)
	err := db.WithContext(ctx).
		Model(&entity.DBTablePost{}).
		Select("id, type, title, created_at").
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Scan(&itemList).Error
	return itemList, err
}
