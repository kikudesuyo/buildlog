package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/external"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QiitaItemFetcher interface {
	FetchAll(context.Context) ([]external.QiitaItem, error)
}

func ImportQiitaItems(ctx context.Context, db *gorm.DB, fetcher QiitaItemFetcher) (int, error) {
	items, err := fetcher.FetchAll(ctx)
	if err != nil {
		return 0, err
	}
	for _, item := range items {
		if err := upsertQiitaItem(ctx, db, item); err != nil {
			return 0, err
		}
	}
	return len(items), nil
}

func upsertQiitaItem(ctx context.Context, db *gorm.DB, item external.QiitaItem) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var source entity.DBTablePostSource
		err := tx.Where("provider = ? AND external_id = ?", external.QiitaProvider, item.ID).First(&source).Error
		if err == nil {
			if err := tx.Model(&entity.DBTablePost{}).Where("id = ?", source.PostID).Updates(map[string]any{
				"title":      item.Title,
				"content":    item.Body,
				"updated_at": item.UpdatedAt,
			}).Error; err != nil {
				return err
			}
			return tx.Model(&source).Updates(map[string]any{"url": item.URL, "updated_at": item.UpdatedAt}).Error
		}
		if err != gorm.ErrRecordNotFound {
			return err
		}

		post := entity.DBTablePost{
			Type:      "tech",
			Title:     item.Title,
			Content:   item.Body,
			Status:    "published",
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		if err := tx.Create(&post).Error; err != nil {
			return err
		}
		return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&entity.DBTablePostSource{
			PostID:     post.ID,
			Provider:   external.QiitaProvider,
			ExternalID: item.ID,
			URL:        item.URL,
		}).Error
	})
}
