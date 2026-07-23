package repository

import (
	"context"
	"fmt"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

func ListDiaryEntries(ctx context.Context, db *gorm.DB) ([]entity.DiaryEntry, error) {
	posts, err := GetPostList(ctx, db)
	if err != nil {
		return nil, err
	}

	entries := make([]entity.DiaryEntry, len(posts))
	for i, post := range posts {
		entries[i] = entity.DiaryEntry{
			ID:       fmt.Sprintf("%d", post.ID),
			Title:    post.Title,
			Excerpt:  post.Content,
			Category: "General",
			Date:     post.CreatedAt.Format("2006年1月2日"),
		}
	}
	return entries, nil
}
