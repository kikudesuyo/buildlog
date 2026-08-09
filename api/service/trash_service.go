package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

// ListDeletedPosts は一覧を取得します。
func ListDeletedPosts(ctx context.Context) ([]entity.DBTablePost, error) {
	return repository.ListDeletedPosts(ctx, database)
}

// RestorePost は削除済みデータを復元します。
func RestorePost(ctx context.Context, id int64) error {
	return repository.RestorePost(ctx, database, id)
}
