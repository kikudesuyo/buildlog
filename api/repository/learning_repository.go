package repository

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

func ListLearnings(ctx context.Context, db *gorm.DB, periodType string, start, end time.Time) ([]entity.DBTableLearning, error) {
	var learnings []entity.DBTableLearning
	err := db.WithContext(ctx).
		Where("period_type = ? AND period_start >= ? AND period_start <= ?", periodType, start, end).
		Order("period_start DESC, id DESC").Find(&learnings).Error
	return learnings, err
}

func CreateLearning(ctx context.Context, db *gorm.DB, learning *entity.DBTableLearning) error {
	return db.WithContext(ctx).Create(learning).Error
}
