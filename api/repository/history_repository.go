package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// GetPostHistory_List は削除されていない投稿履歴を取得します。
func GetPostHistory_List(ctx context.Context, db *gorm.DB) ([]entity.HistoryItem, error) {
	itemList := make([]entity.HistoryItem, 0)
	err := db.WithContext(ctx).
		Raw(`
			SELECT id, type, title, created_at, '' AS url
			FROM posts
			WHERE deleted_at IS NULL
			UNION ALL
			SELECT id, 'tech' AS type, title, published_at AS created_at, url
			FROM external_posts
			ORDER BY created_at DESC
		`).
		Scan(&itemList).Error
	return itemList, err
}
