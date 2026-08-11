package service

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
)

// GetDiary_List は日記一覧を取得し、ページングと閲覧者ごとの反応情報を付与します。
func GetDiary_List(ctx context.Context, all bool, offset int, limit int, sortBy string, sortOrder string, ipAddress string) ([]entity.DBTablePost, error) {
	db := library.GetDB(ctx)
	diary_List, err := repository.GetDiary_List(ctx, db, all, offset, limit, sortBy, sortOrder)
	if err != nil {
		return nil, err
	}
	postID_List := make([]int64, len(diary_List))
	for i := range diary_List {
		postID_List[i] = diary_List[i].ID
	}
	engagement_List, err := repository.GetPostEngagement_List(ctx, db, postID_List, ipAddress)
	if err != nil {
		return nil, err
	}
	for i := range diary_List {
		engagement := engagement_List[diary_List[i].ID]
		diary_List[i].LikesCount = engagement.LikesCount
		diary_List[i].CommentsCount = engagement.CommentsCount
		diary_List[i].HasLiked = engagement.HasLiked
	}
	return diary_List, nil
}

// GetDiaryByID はデータを取得します。
func GetDiaryByID(ctx context.Context, id int64, ipAddress string) (*entity.DBTablePost, error) {
	db := library.GetDB(ctx)
	diary, err := repository.GetDiaryByID(ctx, db, id)
	if err != nil {
		return nil, err
	}
	engagement_List, err := repository.GetPostEngagement_List(ctx, db, []int64{diary.ID}, ipAddress)
	if err != nil {
		return nil, err
	}
	engagement := engagement_List[diary.ID]
	diary.LikesCount = engagement.LikesCount
	diary.CommentsCount = engagement.CommentsCount
	diary.HasLiked = engagement.HasLiked
	return diary, nil
}

// IncrementDiaryViews は公開詳細ページの閲覧数を 1 増やします。
func IncrementDiaryViews(ctx context.Context, id int64) error {
	return repository.IncrementPostViews(ctx, library.GetDB(ctx), "diary", id)
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
