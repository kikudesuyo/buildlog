package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

func ListCommentsByPostID(ctx context.Context, db *gorm.DB, postID int64) ([]entity.DBTableComment, error) {
	commentList := make([]entity.DBTableComment, 0)
	err := db.WithContext(ctx).
		Where("post_id = ?", postID).
		Order("created_at ASC").
		Order("id ASC").
		Find(&commentList).Error
	return commentList, err
}

func CreateComment(ctx context.Context, db *gorm.DB, comment *entity.DBTableComment) error {
	return db.WithContext(ctx).Create(comment).Error
}
