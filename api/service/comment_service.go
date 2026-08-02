package service

import (
	"context"
	"errors"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"github.com/kikudesuyo/buildlog/api/xerror"
	"gorm.io/gorm"
)

func ListCommentsTreeByPostID(ctx context.Context, postID int64) ([]*entity.DBTableComment, error) {
	commentList, err := repository.ListCommentsByPostID(ctx, database, postID)
	if err != nil {
		return nil, xerror.UnknownServerErr(err)
	}

	commentByID := make(map[int64]*entity.DBTableComment, len(commentList))
	rootCommentList := make([]*entity.DBTableComment, 0)
	for i := range commentList {
		comment := &commentList[i]
		comment.Replies = make([]*entity.DBTableComment, 0)
		commentByID[comment.ID] = comment
	}

	for i := range commentList {
		comment := &commentList[i]
		if comment.ParentID == nil {
			rootCommentList = append(rootCommentList, comment)
			continue
		}

		parent, exists := commentByID[*comment.ParentID]
		if !exists || parent.ParentID != nil {
			rootCommentList = append(rootCommentList, comment)
			continue
		}
		parent.Replies = append(parent.Replies, comment)
	}

	return rootCommentList, nil
}

func CreateComment(ctx context.Context, postID int64, req entity.CreateCommentRequest) (*entity.DBTableComment, error) {
	if req.ParentID != nil {
		parent, err := repository.GetCommentByID(ctx, database, *req.ParentID)
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, xerror.ClientResourceNotFoundErr()
		case err != nil:
			return nil, xerror.UnknownServerErr(err)
		case parent.PostID != postID || parent.ParentID != nil:
			return nil, xerror.ClientValidationErr(errors.New("parent comment must be a root comment of the same post"))
		}
	}

	comment := &entity.DBTableComment{
		PostID:   postID,
		ParentID: req.ParentID,
		Content:  req.Content,
	}
	if err := repository.CreateComment(ctx, database, comment); err != nil {
		return nil, xerror.UnknownServerErr(err)
	}

	comment.Replies = make([]*entity.DBTableComment, 0)
	return comment, nil
}
