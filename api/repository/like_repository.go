package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// CreateLike はデータを作成します。
func CreateLike(ctx context.Context, db *gorm.DB, postID int64, ipAddress string) error {
	like := entity.DBTableLike{
		PostID:    postID,
		IPAddress: ipAddress,
	}
	return db.WithContext(ctx).Create(&like).Error
}

// DeleteLike はデータを削除します。
func DeleteLike(ctx context.Context, db *gorm.DB, postID int64, ipAddress string) error {
	return db.WithContext(ctx).
		Where("post_id = ? AND ip_address = ?", postID, ipAddress).
		Delete(&entity.DBTableLike{}).Error
}

// CountLikesByPostID は件数を取得します。
func CountLikesByPostID(ctx context.Context, db *gorm.DB, postID int64) (int64, error) {
	var count int64
	err := db.WithContext(ctx).
		Model(&entity.DBTableLike{}).
		Where("post_id = ?", postID).
		Count(&count).Error
	return count, err
}

// HasLiked は状態を判定します。
func HasLiked(ctx context.Context, db *gorm.DB, postID int64, ipAddress string) (bool, error) {
	var count int64
	err := db.WithContext(ctx).
		Model(&entity.DBTableLike{}).
		Where("post_id = ? AND ip_address = ?", postID, ipAddress).
		Count(&count).Error
	return count > 0, err
}
