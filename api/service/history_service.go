package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
)

func ListPostHistory(ctx context.Context) ([]entity.HistoryItem, error) {
	var items []entity.HistoryItem
	err := database.WithContext(ctx).
		Model(&entity.DBTablePost{}).
		Select("id, type, title, created_at").
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Scan(&items).Error
	return items, err
}
