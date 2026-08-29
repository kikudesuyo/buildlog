package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/external"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

const qiitaUser = "kikudesuyo"

// SyncQiitaArticleList はアプリケーションで利用するQiitaユーザーの記事を同期します。
func SyncQiitaArticleList(ctx context.Context) (int, error) {
	qiitaClient := external.NewQiitaClient(nil, qiitaUser)
	itemList, err := qiitaClient.GetUserArticleList(ctx)
	if err != nil {
		return 0, err
	}
	db := library.GetDB(ctx)
	for _, item := range itemList {
		metadata, _ := qiitaClient.GetOGP(ctx, item.URL)
		if err := syncQiitaArticle(ctx, db, item, metadata); err != nil {
			return 0, err
		}
	}
	return len(itemList), nil
}

func GetTechFeedList(ctx context.Context, db *gorm.DB, all bool, offset, limit int, sortBy, order, ipAddress string) ([]entity.TechFeedItem, error) {
	externalPostList, err := repository.GetExternalPostList(ctx, db, sortBy, order, offset, limit)
	if err != nil {
		return nil, err
	}

	itemList := make([]entity.TechFeedItem, 0, len(externalPostList))
	for _, post := range externalPostList {
		itemList = append(itemList, entity.TechFeedItem{
			Key: fmt.Sprintf("external:%d", post.ID), ID: post.ID, Type: "external", Title: post.Title, Content: post.Excerpt,
			Status: "published", CreatedAt: post.PublishedAt, UpdatedAt: post.UpdatedAt,
			LikesCount: post.LikesCount,
			External:   &entity.ExternalPost{Provider: post.Provider, URL: post.URL, ThumbnailURL: post.ThumbnailURL},
		})
	}

	return itemList, nil
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
			post.LikesCount = item.LikesCount
			post.URL = item.URL
			post.PublishedAt = item.CreatedAt
			post.UpdatedAt = item.UpdatedAt
			return repository.UpdateExternalPost(ctx, tx, post)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		post = &entity.DBTableExternalPost{
			Provider:     external.QiitaProvider,
			ExternalID:   item.ID,
			URL:          item.URL,
			Title:        item.Title,
			Excerpt:      excerptFor(item, metadata),
			ThumbnailURL: metadata.ImageURL,
			LikesCount:   item.LikesCount,
			PublishedAt:  item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
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
