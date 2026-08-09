package service

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
)

// ListDiaries は日記一覧を取得し、ページングと閲覧者ごとの反応情報を付与します。
func ListDiaries(ctx context.Context, all bool, offset int, limit int, ipAddress string) ([]entity.DBTablePost, error) {
	db := library.GetDB(ctx)
	diaryList, err := repository.ListDiaries(ctx, db, all, offset, limit)
	if err != nil {
		return nil, err
	}
	for i := range diaryList {
		count, _ := repository.CountLikesByPostID(ctx, db, diaryList[i].ID)
		commentsCount, _ := repository.CountCommentsByPostID(ctx, db, diaryList[i].ID)
		liked, _ := repository.HasLiked(ctx, db, diaryList[i].ID, ipAddress)
		diaryList[i].LikesCount = count
		diaryList[i].CommentsCount = commentsCount
		diaryList[i].HasLiked = liked
	}
	return diaryList, nil
}

// GetDiaryByID はデータを取得します。
func GetDiaryByID(ctx context.Context, id int64, ipAddress string) (*entity.DBTablePost, error) {
	db := library.GetDB(ctx)
	diary, err := repository.GetDiaryByID(ctx, db, id)
	if err != nil {
		return nil, err
	}
	count, _ := repository.CountLikesByPostID(ctx, db, diary.ID)
	commentsCount, _ := repository.CountCommentsByPostID(ctx, db, diary.ID)
	liked, _ := repository.HasLiked(ctx, db, diary.ID, ipAddress)
	diary.LikesCount = count
	diary.CommentsCount = commentsCount
	diary.HasLiked = liked
	return diary, nil
}

// CreateDiary はデータを作成します。
func CreateDiary(ctx context.Context, req entity.CreateDiaryRequest) (entity.CreateDiaryResponse, error) {
	db := library.GetDB(ctx)
	status := req.Status
	if status == "" {
		status = "draft"
	}
	diary := entity.DBTablePost{
		Title:   req.Title,
		Content: req.Content,
		Status:  status,
	}
	if err := repository.CreateDiary(ctx, db, &diary); err != nil {
		return entity.CreateDiaryResponse{}, err
	}
	return entity.CreateDiaryResponse{
		ID:        diary.ID,
		Title:     diary.Title,
		Content:   diary.Content,
		Status:    diary.Status,
		CreatedAt: diary.CreatedAt.Format(time.RFC3339),
		UpdatedAt: diary.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// UpdateDiary はデータを更新します。
func UpdateDiary(ctx context.Context, id int64, req entity.UpdateDiaryRequest) (entity.UpdateDiaryResponse, error) {
	db := library.GetDB(ctx)
	diary, err := repository.GetDiaryByID(ctx, db, id)
	if err != nil {
		return entity.UpdateDiaryResponse{}, err
	}

	diary.Title = req.Title
	diary.Content = req.Content
	if req.Status != "" {
		diary.Status = req.Status
	}

	if err := repository.UpdateDiary(ctx, db, diary); err != nil {
		return entity.UpdateDiaryResponse{}, err
	}

	return entity.UpdateDiaryResponse{
		ID:        diary.ID,
		Title:     diary.Title,
		Content:   diary.Content,
		Status:    diary.Status,
		CreatedAt: diary.CreatedAt.Format(time.RFC3339),
		UpdatedAt: diary.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// DeleteDiary はデータを削除します。
func DeleteDiary(ctx context.Context, id int64) error {
	db := library.GetDB(ctx)
	return repository.DeleteDiary(ctx, db, id)
}
