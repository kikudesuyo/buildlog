package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

func ListCommentsByPostID(ctx context.Context, postID int64) ([]entity.DBTableComment, error) {
	commentList, err := repository.ListCommentsByPostID(ctx, database, postID)
	if err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	return commentList, nil
}

func CreateComment(ctx context.Context, postID int64, req entity.CreateCommentRequest) (*entity.DBTableComment, error) {
	comment := &entity.DBTableComment{
		PostID:  postID,
		Content: req.Content,
	}
	if err := repository.CreateComment(ctx, database, comment); err != nil {
		return nil, xerror.UnknownServerErr(err)
	}

	return comment, nil
}
