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
	alreadyLiked, err := repository.HasLiked(ctx, databaseFromContext(ctx), postID, ipAddress)
	if err != nil {
		return LikeStatus{}, err
	}
	if !alreadyLiked {
		if err := repository.CreateLike(ctx, databaseFromContext(ctx), postID, ipAddress); err != nil {
			return LikeStatus{}, err
		}
	}

	count, err := repository.CountLikesByPostID(ctx, databaseFromContext(ctx), postID)
	if err != nil {
		return LikeStatus{}, err
	}

	return LikeStatus{
		LikesCount: count,
		HasLiked:   true,
	}, nil
}

// UnlikePost はこの処理に必要な内部処理を実行します。
func UnlikePost(ctx context.Context, postID int64, ipAddress string) (LikeStatus, error) {
	if err := repository.DeleteLike(ctx, databaseFromContext(ctx), postID, ipAddress); err != nil {
		return LikeStatus{}, err
	}

	count, err := repository.CountLikesByPostID(ctx, databaseFromContext(ctx), postID)
	if err != nil {
		return LikeStatus{}, err
	}

	return LikeStatus{
		LikesCount: count,
		HasLiked:   false,
	}, nil
}

// GetLikeStatus はデータを取得します。
func GetLikeStatus(ctx context.Context, postID int64, ipAddress string) (LikeStatus, error) {
	count, err := repository.CountLikesByPostID(ctx, databaseFromContext(ctx), postID)
	if err != nil {
		return LikeStatus{}, err
	}

	hasLiked, err := repository.HasLiked(ctx, databaseFromContext(ctx), postID, ipAddress)
	if err != nil {
		return LikeStatus{}, err
	}

	return LikeStatus{
		LikesCount: count,
		HasLiked:   hasLiked,
	}, nil
}
