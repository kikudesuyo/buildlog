package service

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/external"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

type QiitaItemFetcher interface {
	FetchAll(context.Context) ([]external.QiitaItem, error)
}

type QiitaMetadataFetcher interface {
	FetchOGP(context.Context, string) (external.OGPMetadata, error)
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

func ImportQiitaItems(ctx context.Context, db *gorm.DB, fetcher QiitaItemFetcher, metadataFetcher QiitaMetadataFetcher) (int, error) {
	items, err := fetcher.FetchAll(ctx)
	if err != nil {
		return 0, err
	}
	for _, item := range items {
		metadata := external.OGPMetadata{}
		if metadataFetcher != nil {
			metadata, _ = metadataFetcher.FetchOGP(ctx, item.URL)
		}
		if err := upsertQiitaItem(ctx, db, item, metadata); err != nil {
			return 0, err
		}
	}
	return len(items), nil
}

func upsertQiitaItem(ctx context.Context, db *gorm.DB, item external.QiitaItem, metadata external.OGPMetadata) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var post entity.DBTableExternalPost
		err := tx.Where("provider = ? AND external_id = ?", external.QiitaProvider, item.ID).First(&post).Error
		if err == nil {
			thumbnailURL := post.ThumbnailURL
			if metadata.ImageURL != "" {
				thumbnailURL = metadata.ImageURL
			}
			return tx.Model(&post).Updates(map[string]any{
				"title":         item.Title,
				"excerpt":       excerptFor(item, metadata),
				"thumbnail_url": thumbnailURL,
				"url":           item.URL,
				"published_at":  item.CreatedAt,
				"updated_at":    item.UpdatedAt,
			}).Error
		}
		if err != gorm.ErrRecordNotFound {
			return err
		}

		post = entity.DBTableExternalPost{
			Provider:     external.QiitaProvider,
			ExternalID:   item.ID,
			URL:          item.URL,
			Title:        item.Title,
			Excerpt:      excerptFor(item, metadata),
			ThumbnailURL: metadata.ImageURL,
			PublishedAt:  item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		}
		return tx.Create(&post).Error
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
