package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// GetProfile はデータを取得します。
func GetProfile(ctx context.Context, db *gorm.DB) (*entity.DBTableProfile, error) {
	var profile entity.DBTableProfile
	err := db.WithContext(ctx).First(&profile, 1).Error
	return &profile, err
}

// UpdateProfile はデータを更新します。
func UpdateProfile(ctx context.Context, db *gorm.DB, profile *entity.DBTableProfile) error {
	profile.ID = 1 // シングルトンを保証
	return db.WithContext(ctx).Save(profile).Error
}
