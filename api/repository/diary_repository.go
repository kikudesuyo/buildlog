package repository

import (
	"fmt"

	"github.com/kikudesuyo/buildlog/api/entity"
)

func ListDiaryEntries() []entity.DiaryEntry {
	posts := GetPostList()
	entries := make([]entity.DiaryEntry, len(posts))
	for i, p := range posts {
		entries[i] = entity.DiaryEntry{
			ID:       fmt.Sprintf("%d", p.ID),
			Title:    p.Title,
			Excerpt:  p.Content,
			Category: "General",
			Date:     p.CreatedAt.Format("2006年1月2日"),
		}
	}
	return entries
}
