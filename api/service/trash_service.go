package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
)

func ListDeletedPosts(ctx context.Context) ([]entity.DBTablePost, error) {
	var posts []entity.DBTablePost
	err := database.WithContext(ctx).Unscoped().Where("deleted_at IS NOT NULL").Order("deleted_at DESC").Find(&posts).Error
	return posts, err
}

func RestorePost(ctx context.Context, id int64) error {
	if err := database.WithContext(ctx).Unscoped().Model(&entity.DBTablePost{}).Where("id = ?", id).Update("deleted_at", nil).Error; err != nil {
		return err
	}
	contentCache.Delete("diary:list:false", "diary:list:true", "tech:list:false", "tech:list:true")
	return nil
}
