package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// ListTechs は一覧を取得します。
func ListTechs(ctx context.Context, db *gorm.DB, all bool, offset, limit int) ([]entity.DBTablePost, error) {
	techList := make([]entity.DBTablePost, 0)
	query := db.WithContext(ctx).Where("type = ?", "tech")
	if !all {
		query = query.Where("status = ?", "published")
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	if limit > 0 {
		query = query.Limit(limit)
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

func ListExternalPosts(ctx context.Context, db *gorm.DB) ([]entity.DBTableExternalPost, error) {
	var posts []entity.DBTableExternalPost
	err := db.WithContext(ctx).
		Order("published_at DESC").
		Order("id DESC").
		Find(&posts).Error
	return posts, err
}

// FindExternalPost は外部記事をプロバイダーと外部IDで取得します。
func FindExternalPost(ctx context.Context, db *gorm.DB, provider, externalID string) (*entity.DBTableExternalPost, error) {
	var post entity.DBTableExternalPost
	err := db.WithContext(ctx).Where("provider = ? AND external_id = ?", provider, externalID).First(&post).Error
	return &post, err
}

// InsertExternalPost は外部記事を登録します。
func InsertExternalPost(ctx context.Context, db *gorm.DB, post *entity.DBTableExternalPost) error {
	return db.WithContext(ctx).Create(post).Error
}

// UpdateExternalPost は外部記事を更新します。
func UpdateExternalPost(ctx context.Context, db *gorm.DB, post *entity.DBTableExternalPost) error {
	return db.WithContext(ctx).Save(post).Error
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
