package service

import (
	"context"
	"fmt"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

// ListTechArticles は同期済みの外部技術記事を取得します。
func ListTechArticles(ctx context.Context, offset, limit int) ([]entity.TechFeedItem, error) {
	return ListExternalTechArticles(ctx, library.GetDB(ctx), offset, limit)
}

// ListExternalTechArticles は外部記事専用の一覧を返します。
func ListExternalTechArticles(ctx context.Context, db *gorm.DB, offset, limit int) ([]entity.TechFeedItem, error) {
	externalPosts, err := repository.ListExternalPosts(ctx, db)
	if err != nil {
		return nil, err
	}
	items := make([]entity.TechFeedItem, 0, len(externalPosts))
	for _, post := range externalPosts {
		items = append(items, entity.TechFeedItem{
			Key: fmt.Sprintf("external:%d", post.ID), ID: post.ID, Title: post.Title, Excerpt: post.Excerpt,
			CreatedAt: post.PublishedAt, UpdatedAt: post.UpdatedAt,
			External: entity.ExternalPost{Provider: post.Provider, URL: post.URL, ThumbnailURL: post.ThumbnailURL},
		})
	}
	if offset >= len(items) {
		return []entity.TechFeedItem{}, nil
	}
	end := len(items)
	if limit > 0 && offset+limit < end {
		end = offset + limit
	}
	return items[offset:end], nil
}
