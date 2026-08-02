package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

// ListCommentsByPostID は一覧を取得します。
func ListCommentsByPostID(ctx context.Context, postID int64) ([]entity.DBTableComment, error) {
	commentList, err := repository.ListCommentsByPostID(ctx, databaseFromContext(ctx), postID)
	if err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	return commentList, nil
}

// CreateComment はデータを作成します。
func CreateComment(ctx context.Context, postID int64, req entity.CreateCommentRequest) (*entity.DBTableComment, error) {
	comment := &entity.DBTableComment{
		PostID:  postID,
		Content: req.Content,
	}
	if err := repository.CreateComment(ctx, databaseFromContext(ctx), comment); err != nil {
		return nil, xerror.UnknownServerErr(err)
	}

	return comment, nil
}
