package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

func ListApps(ctx context.Context, db *gorm.DB) ([]entity.DBTableApp, error) {
	apps := make([]entity.DBTableApp, 0)
	err := db.WithContext(ctx).Order("id ASC").Find(&apps).Error
	return apps, err
}

func GetAppByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTableApp, error) {
	var app entity.DBTableApp
	err := db.WithContext(ctx).First(&app, id).Error
	return &app, err
}

func CreateApp(ctx context.Context, db *gorm.DB, app *entity.DBTableApp) error {
	return db.WithContext(ctx).Create(app).Error
}

func UpdateApp(ctx context.Context, db *gorm.DB, app *entity.DBTableApp) error {
	return db.WithContext(ctx).Save(app).Error
}

func DeleteApp(ctx context.Context, db *gorm.DB, id int64) error {
	return db.WithContext(ctx).Delete(&entity.DBTableApp{}, id).Error
}
