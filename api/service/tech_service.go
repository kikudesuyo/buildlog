package service

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
)

// ListTechs は一覧を取得します。
func ListTechs(ctx context.Context, all bool, offset, limit int, ipAddress string) ([]entity.DBTablePost, error) {
	db := library.GetDB(ctx)
	techList, err := repository.ListTechs(ctx, db, all, offset, limit)
	if err != nil {
		return nil, err
	}
	postIDs := make([]int64, len(techList))
	for i := range techList {
		postIDs[i] = techList[i].ID
	}
	engagements, err := repository.GetPostEngagements(ctx, db, postIDs, ipAddress)
	if err != nil {
		return nil, err
	}
	for i := range techList {
		engagement := engagements[techList[i].ID]
		techList[i].LikesCount = engagement.LikesCount
		techList[i].HasLiked = engagement.HasLiked
	}
	return techList, nil
}

// GetTechByID はデータを取得します。
func GetTechByID(ctx context.Context, id int64, ipAddress string) (*entity.DBTablePost, error) {
	db := library.GetDB(ctx)
	tech, err := repository.GetTechByID(ctx, db, id)
	if err != nil {
		return nil, err
	}
	engagements, err := repository.GetPostEngagements(ctx, db, []int64{tech.ID}, ipAddress)
	if err != nil {
		return nil, err
	}
	engagement := engagements[tech.ID]
	tech.LikesCount = engagement.LikesCount
	tech.HasLiked = engagement.HasLiked
	return tech, nil
}

// CreateTech はデータを作成します。
func CreateTech(ctx context.Context, req entity.CreateTechRequest) (entity.CreateTechResponse, error) {
	db := library.GetDB(ctx)
	status := req.Status
	if status == "" {
		status = "draft"
	}
	tech := entity.DBTablePost{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Views:    req.Views,
		Status:   status,
	}

	if err := repository.CreateTech(ctx, db, &tech); err != nil {
		return entity.CreateTechResponse{}, err
	}
	return entity.CreateTechResponse{
		ID:        tech.ID,
		Title:     tech.Title,
		Content:   tech.Content,
		Category:  tech.Category,
		Views:     tech.Views,
		Status:    tech.Status,
		CreatedAt: tech.CreatedAt.Format(time.RFC3339),
		UpdatedAt: tech.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// UpdateTech はデータを更新します。
func UpdateTech(ctx context.Context, id int64, req entity.UpdateTechRequest) (entity.UpdateTechResponse, error) {
	db := library.GetDB(ctx)
	tech, err := repository.GetTechByID(ctx, db, id)
	if err != nil {
		return entity.UpdateTechResponse{}, err
	}

	tech.Title = req.Title
	tech.Content = req.Content
	tech.Category = req.Category
	tech.Views = req.Views
	if req.Status != "" {
		tech.Status = req.Status
	}

	if err := repository.UpdateTech(ctx, db, tech); err != nil {
		return entity.UpdateTechResponse{}, err
	}

	return entity.UpdateTechResponse{
		ID:        tech.ID,
		Title:     tech.Title,
		Content:   tech.Content,
		Category:  tech.Category,
		Views:     tech.Views,
		Status:    tech.Status,
		CreatedAt: tech.CreatedAt.Format(time.RFC3339),
		UpdatedAt: tech.UpdatedAt.Format(time.RFC3339),
	}, nil
}

// DeleteTech はデータを削除します。
func DeleteTech(ctx context.Context, id int64) error {
	db := library.GetDB(ctx)
	return repository.DeleteTech(ctx, db, id)
}
