package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func ListApps(ctx context.Context, db *gorm.DB) ([]entity.DBTableApp, error) {
	return repository.ListApps(ctx, db)
}

func GetAppByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTableApp, error) {
	return repository.GetAppByID(ctx, db, id)
}

func CreateApp(ctx context.Context, db *gorm.DB, req entity.CreateAppRequest) (*entity.DBTableApp, error) {
	app := entity.DBTableApp{
		Slug:        req.Slug,
		Name:        req.Name,
		Category:    req.Category,
		Tags:        req.Tags,
		Description: req.Description,
		Icon:        req.Icon,
		IconURL:     req.IconURL,
		DemoURL:     req.DemoURL,
		CodeURL:     req.CodeURL,
	}

	if err := repository.CreateApp(ctx, db, &app); err != nil {
		return nil, err
	}

	return &app, nil
}

func UpdateApp(ctx context.Context, db *gorm.DB, id int64, req entity.UpdateAppRequest) (*entity.DBTableApp, error) {
	app, err := repository.GetAppByID(ctx, db, id)
	if err != nil {
		return nil, err
	}

	app.Slug = req.Slug
	app.Name = req.Name
	app.Category = req.Category
	app.Tags = req.Tags
	app.Description = req.Description
	app.Icon = req.Icon
	app.IconURL = req.IconURL
	app.DemoURL = req.DemoURL
	app.CodeURL = req.CodeURL

	if err := repository.UpdateApp(ctx, db, app); err != nil {
		return nil, err
	}

	return app, nil
}

func DeleteApp(ctx context.Context, db *gorm.DB, id int64) error {
	return repository.DeleteApp(ctx, db, id)
}
