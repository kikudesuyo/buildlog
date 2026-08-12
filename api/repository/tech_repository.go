package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// GetExternalPostList は外部記事を公開日時順で取得します。
func GetExternalPostList(ctx context.Context, db *gorm.DB, order string, offset, limit int) ([]entity.DBTableExternalPost, error) {
	orderBy := "DESC"
	if order == "asc" {
		orderBy = "ASC"
	}
	var postList []entity.DBTableExternalPost
	query := db.WithContext(ctx).
		Order("published_at " + orderBy).
		Order("id " + orderBy)
	if offset > 0 {
		query = query.Offset(offset)
	}
	if limit > 0 {
		query = query.Limit(limit)
	}
	return postList, query.Find(&postList).Error
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
