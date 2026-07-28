package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

func ListDiaries(ctx context.Context, db *gorm.DB, tag string) ([]entity.DBTablePost, error) {
	diaries := make([]entity.DBTablePost, 0)
	query := db.WithContext(ctx).
		Preload("Tags").
		Where("type = ?", "diary")

	if tag != "" {
		query = query.Joins("JOIN post_tags ON post_tags.post_id = posts.id").
			Joins("JOIN tags ON tags.id = post_tags.tag_id").
			Where("tags.name = ?", tag)
	}

	err := query.Order("created_at DESC").
		Order("id DESC").
		Find(&diaries).Error
	return diaries, err
}

func GetDiaryByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTablePost, error) {
	var diary entity.DBTablePost
	err := db.WithContext(ctx).Preload("Tags").Where("type = ?", "diary").First(&diary, id).Error
	return &diary, err
}

func CreateDiary(ctx context.Context, db *gorm.DB, diary *entity.DBTablePost) error {
	diary.Type = "diary"
	return db.WithContext(ctx).Create(diary).Error
}

func UpdateDiary(ctx context.Context, db *gorm.DB, diary *entity.DBTablePost) error {
	diary.Type = "diary"
	return db.WithContext(ctx).Save(diary).Error
}

func DeleteDiary(ctx context.Context, db *gorm.DB, id int64) error {
	return db.WithContext(ctx).Where("type = ?", "diary").Delete(&entity.DBTablePost{}, id).Error
}
