package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// GetComment_ListByPostID は一覧を取得します。
func GetComment_ListByPostID(ctx context.Context, db *gorm.DB, postID int64) ([]entity.DBTableComment, error) {
	comment_List := make([]entity.DBTableComment, 0)
	err := db.WithContext(ctx).
		Where("post_id = ?", postID).
		Order("created_at ASC").
		Order("id ASC").
		Find(&comment_List).Error
	return comment_List, err
}

// CreateComment はデータを作成します。
func CreateComment(ctx context.Context, db *gorm.DB, comment *entity.DBTableComment) error {
	return db.WithContext(ctx).Create(comment).Error
}

// CountCommentsByPostID は件数を取得します。
func CountCommentsByPostID(ctx context.Context, db *gorm.DB, postID int64) (int64, error) {
	var count int64
	err := db.WithContext(ctx).
		Model(&entity.DBTableComment{}).
		Where("post_id = ?", postID).
		Count(&count).Error
	return count, err
}
