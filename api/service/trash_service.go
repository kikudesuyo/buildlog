package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
)

// ListDeletedPosts は一覧を取得します。
func ListDeletedPosts(ctx context.Context) ([]entity.DBTablePost, error) {
	var posts []entity.DBTablePost
	err := databaseFromContext(ctx).WithContext(ctx).Unscoped().Where("deleted_at IS NOT NULL").Order("deleted_at DESC").Find(&posts).Error
	return posts, err
}

// RestorePost は削除済みデータを復元します。
func RestorePost(ctx context.Context, id int64) error {
	return databaseFromContext(ctx).WithContext(ctx).Unscoped().Model(&entity.DBTablePost{}).Where("id = ?", id).Update("deleted_at", nil).Error
}
