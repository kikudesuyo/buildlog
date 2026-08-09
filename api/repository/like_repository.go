package repository

import (
	"context"
	"fmt"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// PostEngagement は記事に紐づく反応情報を表します。
type PostEngagement struct {
	LikesCount    int64
	CommentsCount int64
	HasLiked      bool
}

// GetPostEngagements は複数記事の反応情報を一括取得します。
func GetPostEngagements(ctx context.Context, db *gorm.DB, postIDs []int64, ipAddress string) (map[int64]PostEngagement, error) {
	engagements := make(map[int64]PostEngagement, len(postIDs))
	if len(postIDs) == 0 {
		return engagements, nil
	}
	for _, postID := range postIDs {
		engagements[postID] = PostEngagement{}
	}

	type count struct {
		PostID int64
		Count  int64
	}
	var likeCounts []count
	if err := db.WithContext(ctx).
		Model(&entity.DBTableLike{}).
		Select("post_id, COUNT(*) AS count").
		Where("post_id IN ?", postIDs).
		Group("post_id").
		Find(&likeCounts).Error; err != nil {
		return nil, fmt.Errorf("list like counts: %w", err)
	}
	for _, item := range likeCounts {
		value := engagements[item.PostID]
		value.LikesCount = item.Count
		engagements[item.PostID] = value
	}

	var commentCounts []count
	if err := db.WithContext(ctx).
		Model(&entity.DBTableComment{}).
		Select("post_id, COUNT(*) AS count").
		Where("post_id IN ?", postIDs).
		Group("post_id").
		Find(&commentCounts).Error; err != nil {
		return nil, fmt.Errorf("list comment counts: %w", err)
	}
	for _, item := range commentCounts {
		value := engagements[item.PostID]
		value.CommentsCount = item.Count
		engagements[item.PostID] = value
	}

	var likedPostIDs []int64
	if err := db.WithContext(ctx).
		Model(&entity.DBTableLike{}).
		Where("post_id IN ? AND ip_address = ?", postIDs, ipAddress).
		Pluck("post_id", &likedPostIDs).Error; err != nil {
		return nil, fmt.Errorf("list liked posts: %w", err)
	}
	for _, postID := range likedPostIDs {
		value := engagements[postID]
		value.HasLiked = true
		engagements[postID] = value
	}

	return engagements, nil
}

// CreateLike はデータを作成します。
func CreateLike(ctx context.Context, db *gorm.DB, postID int64, ipAddress string) error {
	like := entity.DBTableLike{
		PostID:    postID,
		IPAddress: ipAddress,
	}
	return db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(&like).Error
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

// GetLikeStatus は記事のいいね件数と閲覧者の状態を1クエリで取得します。
func GetLikeStatus(ctx context.Context, db *gorm.DB, postID int64, ipAddress string) (int64, bool, error) {
	type status struct {
		LikesCount int64 `gorm:"column:likes_count"`
		HasLiked   bool  `gorm:"column:has_liked"`
	}
	var result status
	err := db.WithContext(ctx).
		Model(&entity.DBTableLike{}).
		Select("COUNT(*) AS likes_count, COALESCE(BOOL_OR(ip_address = ?), false) AS has_liked", ipAddress).
		Where("post_id = ?", postID).
		Scan(&result).Error
	return result.LikesCount, result.HasLiked, err
}
