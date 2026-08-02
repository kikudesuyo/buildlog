package service

import (
	"context"
	"encoding/json"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

// mapToAppResponse はこの処理に必要な内部処理を実行します。
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

// ListApps は一覧を取得します。
func ListApps(ctx context.Context) ([]entity.AppResponse, error) {
	dbApps, err := repository.ListApps(ctx, databaseFromContext(ctx))
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

// GetAppByID はデータを取得します。
func GetAppByID(ctx context.Context, id int64) (*entity.AppResponse, error) {
	dbApp, err := repository.GetAppByID(ctx, databaseFromContext(ctx), id)
	if err != nil {
		return nil, err
	}
	return mapToAppResponse(dbApp)
}

// CreateApp はデータを作成します。
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

	if err := repository.CreateApp(ctx, databaseFromContext(ctx), &app); err != nil {
		return nil, err
	}

	return mapToAppResponse(&app)
}

// UpdateApp はデータを更新します。
func UpdateApp(ctx context.Context, id int64, req entity.UpdateAppRequest) (*entity.AppResponse, error) {
	app, err := repository.GetAppByID(ctx, databaseFromContext(ctx), id)
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

	if err := repository.UpdateApp(ctx, databaseFromContext(ctx), app); err != nil {
		return nil, err
	}

	return mapToAppResponse(app)
}

// DeleteApp はデータを削除します。
func DeleteApp(ctx context.Context, id int64) error {
	return repository.DeleteApp(ctx, databaseFromContext(ctx), id)
}
