package service

import (
	"context"
	"fmt"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func ListTechs(ctx context.Context, db *gorm.DB, all bool, ipAddress string) ([]entity.DBTablePost, error) {
	cacheKey := fmt.Sprintf("tech:list:%t", all)
	var techList []entity.DBTablePost
	if cached, ok := contentCache.Get(cacheKey); ok {
		techList = append([]entity.DBTablePost(nil), cached.([]entity.DBTablePost)...)
	} else {
		var err error
		techList, err = repository.ListTechs(ctx, db, all)
		if err != nil {
			return nil, err
		}
		contentCache.Set(cacheKey, append([]entity.DBTablePost(nil), techList...))
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
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Views:    req.Views,
		Status:   status,
	}
	if err := repository.CreateTech(ctx, db, &tech); err != nil {
		return entity.CreateTechResponse{}, err
	}
	contentCache.Delete("tech:list:false", "tech:list:true")
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

func UpdateTech(ctx context.Context, db *gorm.DB, id int64, req entity.UpdateTechRequest) (entity.UpdateTechResponse, error) {
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
	contentCache.Delete("tech:list:false", "tech:list:true")

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

func DeleteTech(ctx context.Context, db *gorm.DB, id int64) error {
	if err := repository.DeleteTech(ctx, db, id); err != nil {
		return err
	}
	contentCache.Delete("tech:list:false", "tech:list:true")
	return nil
}
