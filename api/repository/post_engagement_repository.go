package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

type postEngagementCount struct {
	PostID int64
	Count  int64
}

type likedPost struct {
	PostID int64
}

// EnrichPostEngagements は投稿一覧のいいね数、コメント数、閲覧者のいいね状態を一括取得します。
func EnrichPostEngagements(ctx context.Context, db *gorm.DB, posts []entity.DBTablePost, ipAddress string) error {
	if len(posts) == 0 {
		return nil
	}

	postIDs := make([]int64, 0, len(posts))
	for _, post := range posts {
		postIDs = append(postIDs, post.ID)
	}

	likeCounts := make([]postEngagementCount, 0)
	if err := db.WithContext(ctx).
		Table("likes").
		Select("post_id, COUNT(*) AS count").
		Where("post_id IN ?", postIDs).
		Group("post_id").
		Scan(&likeCounts).Error; err != nil {
		return err
	}

	commentCounts := make([]postEngagementCount, 0)
	if err := db.WithContext(ctx).
		Table("comments").
		Select("post_id, COUNT(*) AS count").
		Where("post_id IN ?", postIDs).
		Group("post_id").
		Scan(&commentCounts).Error; err != nil {
		return err
	}

	likedPosts := make([]likedPost, 0)
	if err := db.WithContext(ctx).
		Table("likes").
		Select("DISTINCT post_id").
		Where("post_id IN ? AND ip_address = ?", postIDs, ipAddress).
		Scan(&likedPosts).Error; err != nil {
		return err
	}

	likeCountByPostID := make(map[int64]int64, len(likeCounts))
	for _, item := range likeCounts {
		likeCountByPostID[item.PostID] = item.Count
	}
	commentCountByPostID := make(map[int64]int64, len(commentCounts))
	for _, item := range commentCounts {
		commentCountByPostID[item.PostID] = item.Count
	}
	likedPostIDs := make(map[int64]struct{}, len(likedPosts))
	for _, item := range likedPosts {
		likedPostIDs[item.PostID] = struct{}{}
	}

	for i := range posts {
		posts[i].LikesCount = likeCountByPostID[posts[i].ID]
		posts[i].CommentsCount = commentCountByPostID[posts[i].ID]
		_, posts[i].HasLiked = likedPostIDs[posts[i].ID]
	}
	return nil
}
