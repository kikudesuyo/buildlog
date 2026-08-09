package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
)

type LikeStatus struct {
	LikesCount int64 `json:"likes_count"`
	HasLiked   bool  `json:"has_liked"`
}

// LikePost はこの処理に必要な内部処理を実行します。
func LikePost(ctx context.Context, postID int64, ipAddress string) (LikeStatus, error) {
	db := library.GetDB(ctx)
	alreadyLiked, err := repository.HasLiked(ctx, db, postID, ipAddress)
	if err != nil {
		return LikeStatus{}, err
	}
	if !alreadyLiked {
		if err := repository.CreateLike(ctx, db, postID, ipAddress); err != nil {
			return LikeStatus{}, err
		}
	}

	count, err := repository.CountLikesByPostID(ctx, db, postID)
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
	db := library.GetDB(ctx)
	if err := repository.DeleteLike(ctx, db, postID, ipAddress); err != nil {
		return LikeStatus{}, err
	}

	count, err := repository.CountLikesByPostID(ctx, db, postID)
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
	db := library.GetDB(ctx)
	count, err := repository.CountLikesByPostID(ctx, db, postID)
	if err != nil {
		return LikeStatus{}, err
	}

	hasLiked, err := repository.HasLiked(ctx, db, postID, ipAddress)
	if err != nil {
		return LikeStatus{}, err
	}

	return LikeStatus{
		LikesCount: count,
		HasLiked:   hasLiked,
	}, nil
}
