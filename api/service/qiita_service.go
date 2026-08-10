package service

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/external"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

const qiitaUser = "kikudesuyo"

// SyncQiitaArticles はアプリケーションで利用するQiitaユーザーの記事を同期します。
func SyncQiitaArticles(ctx context.Context) (int, error) {
	qiitaClient := external.NewQiitaClient(nil, qiitaUser)
	items, err := qiitaClient.GetUserArticles(ctx)
	if err != nil {
		return 0, err
	}
	db := library.GetDB(ctx)
	for _, item := range items {
		metadata, _ := qiitaClient.GetOGP(ctx, item.URL)
		if err := syncQiitaArticle(ctx, db, item, metadata); err != nil {
			return 0, err
		}
	}
	return len(items), nil
}

func syncQiitaArticle(ctx context.Context, db *gorm.DB, item external.QiitaItem, metadata external.OGPMetadata) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		post, err := repository.FindExternalPost(ctx, tx, external.QiitaProvider, item.ID)
		if err == nil {
			thumbnailURL := post.ThumbnailURL
			if metadata.ImageURL != "" {
				thumbnailURL = metadata.ImageURL
			}
			post.Title = item.Title
			post.Excerpt = excerptFor(item, metadata)
			post.ThumbnailURL = thumbnailURL
			post.URL = item.URL
			post.PublishedAt = item.CreatedAt
			post.UpdatedAt = item.UpdatedAt
			return repository.UpdateExternalPost(ctx, tx, post)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		post = &entity.DBTableExternalPost{
			Provider: external.QiitaProvider, ExternalID: item.ID, URL: item.URL, Title: item.Title,
			Excerpt: excerptFor(item, metadata), ThumbnailURL: metadata.ImageURL,
			PublishedAt: item.CreatedAt, UpdatedAt: item.UpdatedAt,
		}
		return repository.InsertExternalPost(ctx, tx, post)
	})
}

var markdownWhitespace = regexp.MustCompile(`\s+`)

func excerptFor(item external.QiitaItem, metadata external.OGPMetadata) string {
	if metadata.Description != "" {
		return metadata.Description
	}
	content := strings.TrimSpace(markdownWhitespace.ReplaceAllString(item.Body, " "))
	runes := []rune(content)
	if len(runes) > 240 {
		return string(runes[:240])
	}
	return content
}
