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

func ListTechFeed(ctx context.Context, db *gorm.DB, all bool, offset, limit int, ipAddress string) ([]entity.TechFeedItem, error) {
	nativePosts, err := repository.ListTechs(ctx, db, all, 0, 0)
	if err != nil {
		return nil, err
	}
	postIDs := make([]int64, len(nativePosts))
	for i := range nativePosts {
		postIDs[i] = nativePosts[i].ID
	}
	engagements, err := repository.GetPostEngagements(ctx, db, postIDs, ipAddress)
	if err != nil {
		return nil, err
	}
	for i := range nativePosts {
		engagement := engagements[nativePosts[i].ID]
		nativePosts[i].LikesCount = engagement.LikesCount
		nativePosts[i].HasLiked = engagement.HasLiked
	}
	externalPosts, err := repository.ListExternalPosts(ctx, db)
	if err != nil {
		return nil, err
	}

	items := make([]entity.TechFeedItem, 0, len(nativePosts)+len(externalPosts))
	for _, post := range nativePosts {
		items = append(items, entity.TechFeedItem{
			Key: fmt.Sprintf("post:%d", post.ID), ID: post.ID, Type: "post", Title: post.Title, Content: post.Content,
			Views: post.Views, Status: post.Status, CreatedAt: post.CreatedAt, UpdatedAt: post.UpdatedAt,
			LikesCount: post.LikesCount, CommentsCount: post.CommentsCount, HasLiked: post.HasLiked,
		})
	}
	for _, post := range externalPosts {
		items = append(items, entity.TechFeedItem{
			Key: fmt.Sprintf("external:%d", post.ID), ID: post.ID, Type: "external", Title: post.Title, Content: post.Excerpt,
			Status: "published", CreatedAt: post.PublishedAt, UpdatedAt: post.UpdatedAt,
			External: &entity.ExternalPost{Provider: post.Provider, URL: post.URL, ThumbnailURL: post.ThumbnailURL},
		})
	}

	sortTechFeedItems(items)
	if offset >= len(items) {
		return []entity.TechFeedItem{}, nil
	}
	end := len(items)
	if limit > 0 && offset+limit < end {
		end = offset + limit
	}
	return items[offset:end], nil
}

func sortTechFeedItems(items []entity.TechFeedItem) {
	sort.SliceStable(items, func(i, j int) bool {
		if items[i].CreatedAt.Equal(items[j].CreatedAt) {
			return items[i].ID > items[j].ID
		}
		return items[i].CreatedAt.After(items[j].CreatedAt)
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
