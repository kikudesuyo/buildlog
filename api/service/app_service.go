package service

import (
	"context"
	"encoding/json"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

func mapToAppResponse(dbApp *entity.DBTableApp) (*entity.AppResponse, error) {
	var tags []string
	if dbApp.Tags != "" {
		if err := json.Unmarshal([]byte(dbApp.Tags), &tags); err != nil {
			return nil, err
		}
	}
	return &entity.AppResponse{
		ID:          dbApp.ID,
		Slug:        dbApp.Slug,
		Name:        dbApp.Name,
		Category:    dbApp.Category,
		Tags:        tags,
		Description: dbApp.Description,
		Icon:        dbApp.Icon,
		IconURL:     dbApp.IconURL,
		DemoURL:     dbApp.DemoURL,
		CodeURL:     dbApp.CodeURL,
	}, nil
}

func ListApps(ctx context.Context) ([]entity.AppResponse, error) {
	dbApps, err := repository.ListApps(ctx, database)
	if err != nil {
		return nil, err
	}

	appList := make([]entity.AppResponse, 0, len(dbApps))
	for _, dbApp := range dbApps {
		app, err := mapToAppResponse(&dbApp)
		if err != nil {
			return nil, err
		}
		appList = append(appList, *app)
	}

	return appList, nil
}

func GetAppByID(ctx context.Context, id int64) (*entity.AppResponse, error) {
	dbApp, err := repository.GetAppByID(ctx, database, id)
	if err != nil {
		return nil, err
	}
	return mapToAppResponse(dbApp)
}

func CreateApp(ctx context.Context, req entity.CreateAppRequest) (*entity.AppResponse, error) {
	tagsJSON, err := json.Marshal(req.Tags)
	if err != nil {
		return nil, err
	}

	app := entity.DBTableApp{
		Slug:        req.Slug,
		Name:        req.Name,
		Category:    req.Category,
		Tags:        string(tagsJSON),
		Description: req.Description,
		Icon:        req.Icon,
		IconURL:     req.IconURL,
		DemoURL:     req.DemoURL,
		CodeURL:     req.CodeURL,
	}

	if err := repository.CreateApp(ctx, database, &app); err != nil {
		return nil, err
	}
	return mapToAppResponse(&app)
}

func UpdateApp(ctx context.Context, id int64, req entity.UpdateAppRequest) (*entity.AppResponse, error) {
	app, err := repository.GetAppByID(ctx, database, id)
	if err != nil {
		return nil, err
	}

	tagsJSON, err := json.Marshal(req.Tags)
	if err != nil {
		return nil, err
	}

	app.Slug = req.Slug
	app.Name = req.Name
	app.Category = req.Category
	app.Tags = string(tagsJSON)
	app.Description = req.Description
	app.Icon = req.Icon
	app.IconURL = req.IconURL
	app.DemoURL = req.DemoURL
	app.CodeURL = req.CodeURL

	if err := repository.UpdateApp(ctx, database, app); err != nil {
		return nil, err
	}
	return mapToAppResponse(app)
}

func DeleteApp(ctx context.Context, id int64) error {
	if err := repository.DeleteApp(ctx, database, id); err != nil {
		return err
	}
	return nil
}
