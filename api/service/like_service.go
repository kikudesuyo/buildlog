package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/repository"
)

type LikeStatus struct {
	LikesCount int64 `json:"likes_count"`
	HasLiked   bool  `json:"has_liked"`
}

// LikePost はこの処理に必要な内部処理を実行します。
func LikePost(ctx context.Context, postID int64, ipAddress string) (LikeStatus, error) {
	if err := repository.CreateLike(ctx, database, postID, ipAddress); err != nil {
		return LikeStatus{}, err
	}

	count, hasLiked, err := repository.GetLikeStatus(ctx, database, postID, ipAddress)
	if err != nil {
		return LikeStatus{}, err
	}

	return LikeStatus{
		LikesCount: count,
		HasLiked:   hasLiked,
	}, nil
}

// UnlikePost はこの処理に必要な内部処理を実行します。
func UnlikePost(ctx context.Context, postID int64, ipAddress string) (LikeStatus, error) {
	if err := repository.DeleteLike(ctx, database, postID, ipAddress); err != nil {
		return LikeStatus{}, err
	}

	count, hasLiked, err := repository.GetLikeStatus(ctx, database, postID, ipAddress)
	if err != nil {
		return LikeStatus{}, err
	}

	return LikeStatus{
		LikesCount: count,
		HasLiked:   hasLiked,
	}, nil
}

// GetLikeStatus はデータを取得します。
func GetLikeStatus(ctx context.Context, postID int64, ipAddress string) (LikeStatus, error) {
	count, hasLiked, err := repository.GetLikeStatus(ctx, database, postID, ipAddress)
	if err != nil {
		return LikeStatus{}, err
	}

	return LikeStatus{
		LikesCount: count,
		HasLiked:   hasLiked,
	}, nil
}
