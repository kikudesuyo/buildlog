package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// GetAppList は一覧を取得します。
func GetAppList(ctx context.Context, db *gorm.DB) ([]entity.DBTableApp, error) {
	appList := make([]entity.DBTableApp, 0)
	err := db.WithContext(ctx).Order("id ASC").Find(&appList).Error
	return appList, err
}

// GetAppByID はデータを取得します。
func GetAppByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTableApp, error) {
	var app entity.DBTableApp
	err := db.WithContext(ctx).First(&app, id).Error
	return &app, err
}

// CreateApp はデータを作成します。
func CreateApp(ctx context.Context, db *gorm.DB, app *entity.DBTableApp) error {
	return db.WithContext(ctx).Create(app).Error
}

// UpdateApp はデータを更新します。
func UpdateApp(ctx context.Context, db *gorm.DB, app *entity.DBTableApp) error {
	return db.WithContext(ctx).Save(app).Error
}

// DeleteApp はデータを削除します。
func DeleteApp(ctx context.Context, db *gorm.DB, id int64) error {
	return db.WithContext(ctx).Delete(&entity.DBTableApp{}, id).Error
}
