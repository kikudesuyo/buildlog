package repository

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// GetGoalPeriod はデータを取得します。
func GetGoalPeriod(ctx context.Context, db *gorm.DB, periodType string, startsAt time.Time) (*entity.DBTableGoalPeriod, error) {
	var period entity.DBTableGoalPeriod
	err := db.WithContext(ctx).
		Where("period_type = ? AND starts_at = ?", periodType, startsAt).
		Preload("Goals", func(db *gorm.DB) *gorm.DB { return db.Order("id ASC") }).
		First(&period).Error
	return &period, err
}

// GetGoalPeriodList は期間別の目標履歴を取得します。
func GetGoalPeriodList(ctx context.Context, db *gorm.DB, periodType string, before time.Time) ([]entity.DBTableGoalPeriod, error) {
	periodList := make([]entity.DBTableGoalPeriod, 0)
	err := db.WithContext(ctx).
		Where("period_type = ? AND starts_at < ?", periodType, before).
		Preload("Goals", func(db *gorm.DB) *gorm.DB { return db.Order("id ASC") }).
		Order("starts_at DESC").
		Find(&periodList).Error
	return periodList, err
}

// SaveGoalPeriod はデータを保存します。
func SaveGoalPeriod(ctx context.Context, db *gorm.DB, period *entity.DBTableGoalPeriod) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("period_id = ?", period.ID).Delete(&entity.DBTableGoal{}).Error; err != nil {
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
