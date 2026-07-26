package service

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func ListDiaries(ctx context.Context, db *gorm.DB) ([]entity.DBTablePost, error) {
	return repository.ListDiaries(ctx, db)
}

func GetDiaryByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTablePost, error) {
	return repository.GetDiaryByID(ctx, db, id)
}

func CreateDiary(ctx context.Context, db *gorm.DB, title, content string) (entity.CreateDiaryResponse, error) {
	diary := entity.DBTablePost{
		Title:   title,
		Content: content,
	}
	if err := repository.CreateDiary(ctx, db, &diary); err != nil {
		return entity.CreateDiaryResponse{}, err
	}
	return entity.CreateDiaryResponse{
		ID:        diary.ID,
		Title:     diary.Title,
		Content:   diary.Content,
		CreatedAt: diary.CreatedAt.Format(time.RFC3339),
		UpdatedAt: diary.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func UpdateDiary(ctx context.Context, db *gorm.DB, id int64, title, content string) (entity.UpdateDiaryResponse, error) {
	diary, err := repository.GetDiaryByID(ctx, db, id)
	if err != nil {
		return entity.UpdateDiaryResponse{}, err
	}

	diary.Title = title
	diary.Content = content

	if err := repository.UpdateDiary(ctx, db, diary); err != nil {
		return entity.UpdateDiaryResponse{}, err
	}

	return entity.UpdateDiaryResponse{
		ID:        diary.ID,
		Title:     diary.Title,
		Content:   diary.Content,
		CreatedAt: diary.CreatedAt.Format(time.RFC3339),
		UpdatedAt: diary.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func DeleteDiary(ctx context.Context, db *gorm.DB, id int64) error {
	return repository.DeleteDiary(ctx, db, id)
}
