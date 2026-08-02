package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// ListTechs は一覧を取得します。
func ListTechs(ctx context.Context, db *gorm.DB, all bool, queryText string) ([]entity.DBTablePost, error) {
	techList := make([]entity.DBTablePost, 0)
	query := db.WithContext(ctx).Where("type = ?", "tech")
	if !all {
		query = query.Where("status = ?", "published")
	}
	if queryText != "" {
		pattern := "%" + queryText + "%"
		query = query.Where("title ILIKE ? OR content ILIKE ?", pattern, pattern)
	}
	err := query.
		Order("created_at DESC").
		Order("id DESC").
		Find(&techList).Error
	return techList, err
}

// GetTechByID はデータを取得します。
func GetTechByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTablePost, error) {
	var tech entity.DBTablePost
	err := db.WithContext(ctx).Where("type = ?", "tech").First(&tech, id).Error
	return &tech, err
}

// CreateTech はデータを作成します。
func CreateTech(ctx context.Context, db *gorm.DB, tech *entity.DBTablePost) error {
	tech.Type = "tech"
	return db.WithContext(ctx).Create(tech).Error
}

// UpdateTech はデータを更新します。
func UpdateTech(ctx context.Context, db *gorm.DB, tech *entity.DBTablePost) error {
	tech.Type = "tech"
	return db.WithContext(ctx).Save(tech).Error
}

// DeleteTech はデータを削除します。
func DeleteTech(ctx context.Context, db *gorm.DB, id int64) error {
	return db.WithContext(ctx).Where("type = ?", "tech").Delete(&entity.DBTablePost{}, id).Error
}
