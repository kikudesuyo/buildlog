package service

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func ListTechs(ctx context.Context, db *gorm.DB) ([]entity.DBTablePost, error) {
	return repository.ListTechs(ctx, db)
}

func GetTechByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTablePost, error) {
	return repository.GetTechByID(ctx, db, id)
}

func CreateTech(ctx context.Context, db *gorm.DB, req entity.CreateTechRequest) (entity.CreateTechResponse, error) {
	tech := entity.DBTablePost{
		Title:        req.Title,
		Excerpt:      req.Excerpt,
		Category:     req.Category,
		ReadTime:     req.ReadTime,
		Views:        req.Views,
		IsNewsletter: req.IsNewsletter,
	}
	if err := repository.CreateTech(ctx, db, &tech); err != nil {
		return entity.CreateTechResponse{}, err
	}
	return entity.CreateTechResponse{
		ID:           tech.ID,
		Title:        tech.Title,
		Excerpt:      tech.Excerpt,
		Category:     tech.Category,
		ReadTime:     tech.ReadTime,
		Views:        tech.Views,
		IsNewsletter: tech.IsNewsletter,
		CreatedAt:    tech.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    tech.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func UpdateTech(ctx context.Context, db *gorm.DB, id int64, req entity.UpdateTechRequest) (entity.UpdateTechResponse, error) {
	tech, err := repository.GetTechByID(ctx, db, id)
	if err != nil {
		return entity.UpdateTechResponse{}, err
	}

	tech.Title = req.Title
	tech.Excerpt = req.Excerpt
	tech.Category = req.Category
	tech.ReadTime = req.ReadTime
	tech.Views = req.Views
	tech.IsNewsletter = req.IsNewsletter

	if err := repository.UpdateTech(ctx, db, tech); err != nil {
		return entity.UpdateTechResponse{}, err
	}

	return entity.UpdateTechResponse{
		ID:           tech.ID,
		Title:        tech.Title,
		Excerpt:      tech.Excerpt,
		Category:     tech.Category,
		ReadTime:     tech.ReadTime,
		Views:        tech.Views,
		IsNewsletter: tech.IsNewsletter,
		CreatedAt:    tech.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    tech.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func DeleteTech(ctx context.Context, db *gorm.DB, id int64) error {
	return repository.DeleteTech(ctx, db, id)
}
