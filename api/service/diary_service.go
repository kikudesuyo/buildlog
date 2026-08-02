package service

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

// ListDiaries は日記一覧を取得し、ページングと閲覧者ごとの反応情報を付与します。
func ListDiaries(ctx context.Context, all bool, offset int, limit int, ipAddress string) ([]entity.DBTablePost, error) {
	diaryList, err := repository.ListDiaries(ctx, database, all, offset, limit)
	if err != nil {
		return nil, err
	}
	if err := repository.EnrichPostEngagements(ctx, database, diaryList, ipAddress); err != nil {
		return nil, err
	}
	return diaryList, nil
}

// GetDiaryByID はデータを取得します。
func GetDiaryByID(ctx context.Context, id int64, ipAddress string) (*entity.DBTablePost, error) {
	diary, err := repository.GetDiaryByID(ctx, database, id)
	if err != nil {
		return nil, err
	}
	posts := []entity.DBTablePost{*diary}
	if err := repository.EnrichPostEngagements(ctx, database, posts, ipAddress); err != nil {
		return nil, err
	}
	*diary = posts[0]
	return diary, nil
}

// CreateDiary はデータを作成します。
func CreateDiary(ctx context.Context, req entity.CreateDiaryRequest) (entity.CreateDiaryResponse, error) {
	status := req.Status
	if status == "" {
		status = "draft"
	}
	diary := entity.DBTablePost{
		Title:   req.Title,
		Content: req.Content,
		Status:  status,
	}
	if err := repository.CreateDiary(ctx, database, &diary); err != nil {
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
	diary, err := repository.GetDiaryByID(ctx, database, id)
	if err != nil {
		return entity.UpdateDiaryResponse{}, err
	}

	diary.Title = req.Title
	diary.Content = req.Content
	if req.Status != "" {
		diary.Status = req.Status
	}

	if err := repository.UpdateDiary(ctx, database, diary); err != nil {
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
	return repository.DeleteDiary(ctx, database, id)
}
