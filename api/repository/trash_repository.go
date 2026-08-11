package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// GetDeletedPost_List は一覧を取得します。
func GetDeletedPost_List(ctx context.Context, db *gorm.DB) ([]entity.DBTablePost, error) {
	var post_List []entity.DBTablePost
	err := db.WithContext(ctx).
		Unscoped().
		Where("deleted_at IS NOT NULL").
		Order("deleted_at DESC").
		Find(&post_List).Error
	return post_List, err
}

// RestorePost は削除済みデータを復元します。
func RestorePost(ctx context.Context, db *gorm.DB, id int64) error {
	return db.WithContext(ctx).
		Unscoped().
		Model(&entity.DBTablePost{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}
