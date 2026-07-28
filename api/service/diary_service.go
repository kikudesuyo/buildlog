package service

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func ListDiaries(ctx context.Context, db *gorm.DB, all bool, ipAddress string) ([]entity.DBTablePost, error) {
	diaries, err := repository.ListDiaries(ctx, db, all)
	if err != nil {
		return nil, err
	}
	for i := range diaries {
		count, _ := repository.CountLikesByPostID(ctx, db, diaries[i].ID)
		liked, _ := repository.HasLiked(ctx, db, diaries[i].ID, ipAddress)
		diaries[i].LikesCount = count
		diaries[i].HasLiked = liked
	}
	return diaries, nil
}

func GetDiaryByID(ctx context.Context, db *gorm.DB, id int64, ipAddress string) (*entity.DBTablePost, error) {
	diary, err := repository.GetDiaryByID(ctx, db, id)
	if err != nil {
		return nil, err
	}
	count, _ := repository.CountLikesByPostID(ctx, db, diary.ID)
	liked, _ := repository.HasLiked(ctx, db, diary.ID, ipAddress)
	diary.LikesCount = count
	diary.HasLiked = liked
	return diary, nil
}

func CreateDiary(ctx context.Context, db *gorm.DB, req entity.CreateDiaryRequest) (entity.CreateDiaryResponse, error) {
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

func UpdateDiary(ctx context.Context, db *gorm.DB, id int64, req entity.UpdateDiaryRequest) (entity.UpdateDiaryResponse, error) {
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

func DeleteDiary(ctx context.Context, db *gorm.DB, id int64) error {
	return repository.DeleteDiary(ctx, db, id)
}
