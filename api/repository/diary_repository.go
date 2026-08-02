package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

func ListDiaries(ctx context.Context, db *gorm.DB, all bool, offset int, limit int) ([]entity.DBTablePost, error) {
	diaryList := make([]entity.DBTablePost, 0)
	query := db.WithContext(ctx).Where("type = ?", "diary")
	if !all {
		query = query.Where("status = ?", "published")
	}
	query = query.Order("created_at DESC").Order("id DESC")
	if !all {
		query = query.Offset(offset)
		if limit > 0 {
			query = query.Limit(limit)
		}
	}
	err := query.Find(&diaryList).Error
	return diaryList, err
}

func GetDiaryByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTablePost, error) {
	var diary entity.DBTablePost
	err := db.WithContext(ctx).Where("type = ?", "diary").First(&diary, id).Error
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
