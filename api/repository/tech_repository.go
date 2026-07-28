package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

func ListTechs(ctx context.Context, db *gorm.DB, tag string) ([]entity.DBTablePost, error) {
	techs := make([]entity.DBTablePost, 0)
	query := db.WithContext(ctx).
		Preload("Tags").
		Where("type = ?", "tech")

	if tag != "" {
		query = query.Joins("JOIN post_tags ON post_tags.post_id = posts.id").
			Joins("JOIN tags ON tags.id = post_tags.tag_id").
			Where("tags.name = ?", tag)
	}

	err := query.Order("created_at DESC").
		Order("id DESC").
		Find(&techs).Error
	return techs, err
}

func GetTechByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTablePost, error) {
	var tech entity.DBTablePost
	err := db.WithContext(ctx).Preload("Tags").Where("type = ?", "tech").First(&tech, id).Error
	return &tech, err
}

func CreateTech(ctx context.Context, db *gorm.DB, tech *entity.DBTablePost) error {
	tech.Type = "tech"
	return db.WithContext(ctx).Create(tech).Error
}

func UpdateTech(ctx context.Context, db *gorm.DB, tech *entity.DBTablePost) error {
	tech.Type = "tech"
	return db.WithContext(ctx).Save(tech).Error
}

func DeleteTech(ctx context.Context, db *gorm.DB, id int64) error {
	return db.WithContext(ctx).Where("type = ?", "tech").Delete(&entity.DBTablePost{}, id).Error
}
