package repository

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

func GetGoalPeriod(ctx context.Context, db *gorm.DB, periodType string, startsAt time.Time) (*entity.DBTableGoalPeriod, error) {
	var period entity.DBTableGoalPeriod
	err := db.WithContext(ctx).
		Where("period_type = ? AND starts_at = ?", periodType, startsAt).
		Preload("Goals", func(db *gorm.DB) *gorm.DB { return db.Order("id ASC") }).
		First(&period).Error
	return &period, err
}

func SaveGoalPeriod(ctx context.Context, db *gorm.DB, period *entity.DBTableGoalPeriod) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", period.ID).Delete(&entity.DBTableGoal{}).Error; err != nil {
			return err
		}
		if err := tx.Omit("Goals").Save(period).Error; err != nil {
			return err
		}
		for i := range period.Goals {
			period.Goals[i].PeriodID = period.ID
			if err := tx.Create(&period.Goals[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
