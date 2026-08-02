package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

func ListDeletedPosts(ctx context.Context, db *gorm.DB) ([]entity.DBTablePost, error) {
	var posts []entity.DBTablePost
	err := db.WithContext(ctx).
		Unscoped().
		Where("deleted_at IS NOT NULL").
		Order("deleted_at DESC").
		Find(&posts).Error
	return posts, err
}

func RestorePost(ctx context.Context, db *gorm.DB, id int64) error {
	return db.WithContext(ctx).
		Unscoped().
		Model(&entity.DBTablePost{}).
		Where("id = ?", id).
		Update("deleted_at", nil).Error
}
