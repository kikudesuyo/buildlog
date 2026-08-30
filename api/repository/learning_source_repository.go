package repository

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// FetchPublishedPostsForLearning は指定期間の公開済みアプリ投稿を取得します。
func FetchPublishedPostsForLearning(ctx context.Context, db *gorm.DB, start, end time.Time) ([]entity.DBTablePost, error) {
	var posts []entity.DBTablePost
	err := db.WithContext(ctx).
		Where("status = ? AND created_at >= ? AND created_at < ?", "published", start, end.AddDate(0, 0, 1)).
		Order("created_at ASC, id ASC").Find(&posts).Error
	return posts, err
}

// FetchExternalPostsForLearning は指定期間に公開された外部投稿を取得します。
func FetchExternalPostsForLearning(ctx context.Context, db *gorm.DB, start, end time.Time) ([]entity.DBTableExternalPost, error) {
	var posts []entity.DBTableExternalPost
	err := db.WithContext(ctx).
		Where("published_at >= ? AND published_at < ?", start, end.AddDate(0, 0, 1)).
		Order("published_at ASC, id ASC").Find(&posts).Error
	return posts, err
}
