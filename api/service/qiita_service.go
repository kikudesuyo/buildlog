package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/external"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

const qiitaUser = "kikudesuyo"

// SyncQiitaArticle_List はアプリケーションで利用するQiitaユーザーの記事を同期します。
func SyncQiitaArticle_List(ctx context.Context) (int, error) {
	qiitaClient := external.NewQiitaClient(nil, qiitaUser)
	item_List, err := qiitaClient.GetUserArticle_List(ctx)
	if err != nil {
		return 0, err
	}
	db := library.GetDB(ctx)
	for _, item := range item_List {
		metadata, _ := qiitaClient.GetOGP(ctx, item.URL)
		if err := syncQiitaArticle(ctx, db, item, metadata); err != nil {
			return 0, err
		}
	}
	return len(item_List), nil
}

func GetTechFeed_List(ctx context.Context, db *gorm.DB, all bool, offset, limit int, order, ipAddress string) ([]entity.TechFeedItem, error) {
	externalPost_List, err := repository.GetExternalPost_List(ctx, db, order)
	if err != nil {
		return nil, err
	}

	item_List := make([]entity.TechFeedItem, 0, len(externalPost_List))
	for _, post := range externalPost_List {
		item_List = append(item_List, entity.TechFeedItem{
			Key: fmt.Sprintf("external:%d", post.ID), ID: post.ID, Type: "external", Title: post.Title, Content: post.Excerpt,
			Status: "published", CreatedAt: post.PublishedAt, UpdatedAt: post.UpdatedAt,
			LikesCount: post.LikesCount,
			External:   &entity.ExternalPost{Provider: post.Provider, URL: post.URL, ThumbnailURL: post.ThumbnailURL},
		})
	}

	sortTechFeed_List(item_List, order)
	if offset >= len(item_List) {
		return []entity.TechFeedItem{}, nil
	}
	end := len(item_List)
	if limit > 0 && offset+limit < end {
		end = offset + limit
	}
	return item_List[offset:end], nil
}

func sortTechFeed_List(item_List []entity.TechFeedItem, order string) {
	sort.SliceStable(item_List, func(i, j int) bool {
		if item_List[i].CreatedAt.Equal(item_List[j].CreatedAt) {
			if order == "asc" {
				return item_List[i].ID < item_List[j].ID
			}
			return item_List[i].ID > item_List[j].ID
		}
		if order == "asc" {
			return item_List[i].CreatedAt.Before(item_List[j].CreatedAt)
		}
		return item_List[i].CreatedAt.After(item_List[j].CreatedAt)
	})
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
