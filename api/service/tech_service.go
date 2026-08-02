package service

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

func ListTechs(ctx context.Context, all bool, ipAddress string) ([]entity.DBTablePost, error) {
	techList, err := repository.ListTechs(ctx, database, all)
	if err != nil {
		return nil, err
	}
	for i := range techList {
		count, _ := repository.CountLikesByPostID(ctx, database, techList[i].ID)
		liked, _ := repository.HasLiked(ctx, database, techList[i].ID, ipAddress)
		techList[i].LikesCount = count
		techList[i].HasLiked = liked
	}
	return techList, nil
}

func GetTechByID(ctx context.Context, id int64, ipAddress string) (*entity.DBTablePost, error) {
	tech, err := repository.GetTechByID(ctx, database, id)
	if err != nil {
		return nil, err
	}
	count, _ := repository.CountLikesByPostID(ctx, database, tech.ID)
	liked, _ := repository.HasLiked(ctx, database, tech.ID, ipAddress)
	tech.LikesCount = count
	tech.HasLiked = liked
	return tech, nil
}

func CreateTech(ctx context.Context, req entity.CreateTechRequest) (entity.CreateTechResponse, error) {
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
	if err := repository.CreateTech(ctx, database, &tech); err != nil {
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

func UpdateTech(ctx context.Context, id int64, req entity.UpdateTechRequest) (entity.UpdateTechResponse, error) {
	tech, err := repository.GetTechByID(ctx, database, id)
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

	if err := repository.UpdateTech(ctx, database, tech); err != nil {
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

func DeleteTech(ctx context.Context, id int64) error {
	return repository.DeleteTech(ctx, database, id)
}
