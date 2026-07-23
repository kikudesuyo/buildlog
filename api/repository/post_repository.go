package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

func GetPostList(ctx context.Context, db *gorm.DB) ([]entity.DBTablePost, error) {
	posts := make([]entity.DBTablePost, 0)
	err := db.WithContext(ctx).
		Order("created_at DESC").
		Order("id DESC").
		Find(&posts).Error
	return posts, err
}
