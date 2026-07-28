package service

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func ListTechs(ctx context.Context, db *gorm.DB, tag string) ([]entity.DBTablePost, error) {
	return repository.ListTechs(ctx, db, tag)
}

func GetTechByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTablePost, error) {
	return repository.GetTechByID(ctx, db, id)
}

func CreateTech(ctx context.Context, db *gorm.DB, req entity.CreateTechRequest) (entity.CreateTechResponse, error) {
	tech := entity.DBTablePost{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Views:    req.Views,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := repository.CreateTech(ctx, tx, &tech); err != nil {
			return err
		}
		if err := savePostTags(tx, &tech, req.Tags); err != nil {
			return err
		}
		return tx.Preload("Tags").First(&tech, tech.ID).Error
	})
	if err != nil {
		return entity.CreateTechResponse{}, err
	}

	return entity.CreateTechResponse{
		ID:        tech.ID,
		Title:     tech.Title,
		Content:   tech.Content,
		Category:  tech.Category,
		Views:     tech.Views,
		Tags:      mapTagsToStrings(tech.Tags),
		CreatedAt: tech.CreatedAt.Format(time.RFC3339),
		UpdatedAt: tech.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func UpdateTech(ctx context.Context, db *gorm.DB, id int64, req entity.UpdateTechRequest) (entity.UpdateTechResponse, error) {
	var tech *entity.DBTablePost
	err := db.Transaction(func(tx *gorm.DB) error {
		var err error
		tech, err = repository.GetTechByID(ctx, tx, id)
		if err != nil {
			return err
		}

		tech.Title = req.Title
		tech.Content = req.Content
		tech.Category = req.Category
		tech.Views = req.Views

		if err := repository.UpdateTech(ctx, tx, tech); err != nil {
			return err
		}
		if err := savePostTags(tx, tech, req.Tags); err != nil {
			return err
		}
		return tx.Preload("Tags").First(tech, tech.ID).Error
	})
	if err != nil {
		return entity.UpdateTechResponse{}, err
	}

	return entity.UpdateTechResponse{
		ID:        tech.ID,
		Title:     tech.Title,
		Content:   tech.Content,
		Category:  tech.Category,
		Views:     tech.Views,
		Tags:      mapTagsToStrings(tech.Tags),
		CreatedAt: tech.CreatedAt.Format(time.RFC3339),
		UpdatedAt: tech.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func DeleteTech(ctx context.Context, db *gorm.DB, id int64) error {
	return repository.DeleteTech(ctx, db, id)
}
