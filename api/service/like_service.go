package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/repository"
)

type LikeStatus struct {
	LikesCount int64 `json:"likes_count"`
	HasLiked   bool  `json:"has_liked"`
}

func LikePost(ctx context.Context, postID int64, ipAddress string) (LikeStatus, error) {
	alreadyLiked, err := repository.HasLiked(ctx, database, postID, ipAddress)
	if err != nil {
		return LikeStatus{}, err
	}
	if !alreadyLiked {
		if err := repository.CreateLike(ctx, database, postID, ipAddress); err != nil {
			return LikeStatus{}, err
		}
	}

	count, err := repository.CountLikesByPostID(ctx, database, postID)
	if err != nil {
		return LikeStatus{}, err
	}

	return LikeStatus{
		LikesCount: count,
		HasLiked:   true,
	}, nil
}

func UnlikePost(ctx context.Context, postID int64, ipAddress string) (LikeStatus, error) {
	if err := repository.DeleteLike(ctx, database, postID, ipAddress); err != nil {
		return LikeStatus{}, err
	}

	count, err := repository.CountLikesByPostID(ctx, database, postID)
	if err != nil {
		return LikeStatus{}, err
	}

	return LikeStatus{
		LikesCount: count,
		HasLiked:   false,
	}, nil
}

func GetLikeStatus(ctx context.Context, postID int64, ipAddress string) (LikeStatus, error) {
	count, err := repository.CountLikesByPostID(ctx, database, postID)
	if err != nil {
		return LikeStatus{}, err
	}

	hasLiked, err := repository.HasLiked(ctx, database, postID, ipAddress)
	if err != nil {
		return LikeStatus{}, err
	}

	return LikeStatus{
		LikesCount: count,
		HasLiked:   hasLiked,
	}, nil
}
