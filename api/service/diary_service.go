package service

import (
	"context"
	"fmt"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func ListDiaries(ctx context.Context, db *gorm.DB, all bool, ipAddress string) ([]entity.DBTablePost, error) {
	cacheKey := fmt.Sprintf("diary:list:%t", all)
	var diaryList []entity.DBTablePost
	if cached, ok := contentCache.Get(cacheKey); ok {
		diaryList = append([]entity.DBTablePost(nil), cached.([]entity.DBTablePost)...)
	} else {
		var err error
		diaryList, err = repository.ListDiaries(ctx, db, all)
		if err != nil {
			return nil, err
		}
		contentCache.Set(cacheKey, append([]entity.DBTablePost(nil), diaryList...))
	}
	for i := range diaryList {
		count, _ := repository.CountLikesByPostID(ctx, db, diaryList[i].ID)
		liked, _ := repository.HasLiked(ctx, db, diaryList[i].ID, ipAddress)
		diaryList[i].LikesCount = count
		diaryList[i].HasLiked = liked
	}
	return diaryList, nil
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
	contentCache.Delete("diary:list:false", "diary:list:true")
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
	contentCache.Delete("diary:list:false", "diary:list:true")

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
	if err := repository.DeleteDiary(ctx, db, id); err != nil {
		return err
	}
	contentCache.Delete("diary:list:false", "diary:list:true")
	return nil
}
