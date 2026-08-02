package repository

import (
	"context"
	"strings"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// ReplacePostTags resolves tag names and replaces the post's tag associations.
func ReplacePostTags(ctx context.Context, db *gorm.DB, postID int64, names []string) error {
	tags := make([]entity.DBTableTag, 0, len(names))
	seen := make(map[string]struct{}, len(names))
	for _, name := range names {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		if _, ok := seen[name]; ok {
			continue
		}
		seen[name] = struct{}{}

		var tag entity.DBTableTag
		if err := db.WithContext(ctx).
			Where("name = ?", name).
			FirstOrCreate(&tag, entity.DBTableTag{Name: name}).Error; err != nil {
			return err
		}
		tags = append(tags, tag)
	}

	post := entity.DBTablePost{ID: postID}
	return db.WithContext(ctx).Model(&post).Association("Tags").Replace(tags)
}
