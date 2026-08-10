package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

func ListExternalPosts(ctx context.Context, db *gorm.DB) ([]entity.DBTableExternalPost, error) {
	var posts []entity.DBTableExternalPost
	err := db.WithContext(ctx).
		Order("published_at DESC").
		Order("id DESC").
		Find(&posts).Error
	return posts, err
}

// FindExternalPost は外部記事をプロバイダーと外部IDで取得します。
func FindExternalPost(ctx context.Context, db *gorm.DB, provider, externalID string) (*entity.DBTableExternalPost, error) {
	var post entity.DBTableExternalPost
	err := db.WithContext(ctx).Where("provider = ? AND external_id = ?", provider, externalID).First(&post).Error
	return &post, err
}

// InsertExternalPost は外部記事を登録します。
func InsertExternalPost(ctx context.Context, db *gorm.DB, post *entity.DBTableExternalPost) error {
	return db.WithContext(ctx).Create(post).Error
}

// UpdateExternalPost は外部記事を更新します。
func UpdateExternalPost(ctx context.Context, db *gorm.DB, post *entity.DBTableExternalPost) error {
	return db.WithContext(ctx).Save(post).Error
}
