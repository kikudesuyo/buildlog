package service

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func ListTechs(ctx context.Context, db *gorm.DB, all bool, ipAddress string, tag string) ([]entity.DBTablePost, error) {
	techList, err := repository.ListTechs(ctx, db, all, tag)
	if err != nil {
		return nil, err
	}
	for i := range techList {
		count, _ := repository.CountLikesByPostID(ctx, db, techList[i].ID)
		liked, _ := repository.HasLiked(ctx, db, techList[i].ID, ipAddress)
		techList[i].LikesCount = count
		techList[i].HasLiked = liked
	}
	return techList, nil
}

func GetTechByID(ctx context.Context, db *gorm.DB, id int64, ipAddress string) (*entity.DBTablePost, error) {
	tech, err := repository.GetTechByID(ctx, db, id)
	if err != nil {
		return nil, err
	}
	count, _ := repository.CountLikesByPostID(ctx, db, tech.ID)
	liked, _ := repository.HasLiked(ctx, db, tech.ID, ipAddress)
	tech.LikesCount = count
	tech.HasLiked = liked
	return tech, nil
}

func CreateTech(ctx context.Context, db *gorm.DB, req entity.CreateTechRequest) (entity.CreateTechResponse, error) {
	status := req.Status
	if status == "" {
		status = "draft"
	}
	tech := entity.DBTablePost{
		Title: req.Title, Content: req.Content, Category: req.Category, Views: req.Views, Status: status,
	}

	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := repository.CreateTech(ctx, tx, &tech); err != nil {
			return err
		}
		if err := repository.ReplacePostTags(ctx, tx, tech.ID, req.Tags); err != nil {
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
		Status:    tech.Status,
		Tags:      mapTagsToStrings(tech.Tags),
		CreatedAt: tech.CreatedAt.Format(time.RFC3339),
		UpdatedAt: tech.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func UpdateTech(ctx context.Context, db *gorm.DB, id int64, req entity.UpdateTechRequest) (entity.UpdateTechResponse, error) {
	var tech *entity.DBTablePost
	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		tech, err = repository.GetTechByID(ctx, tx, id)
		if err != nil {
			return err
		}
		tech.Title = req.Title
		tech.Content = req.Content
		tech.Category = req.Category
		tech.Views = req.Views
		if req.Status != "" {
			tech.Status = req.Status
		}
		if err := repository.UpdateTech(ctx, tx, tech); err != nil {
			return err
		}
		if err := repository.ReplacePostTags(ctx, tx, tech.ID, req.Tags); err != nil {
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
		Status:    tech.Status,
		Tags:      mapTagsToStrings(tech.Tags),
		CreatedAt: tech.CreatedAt.Format(time.RFC3339),
		UpdatedAt: tech.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func DeleteTech(ctx context.Context, db *gorm.DB, id int64) error {
	return repository.DeleteTech(ctx, db, id)
}
