package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
)

// GetDeletedPost_List は一覧を取得します。
func GetDeletedPost_List(ctx context.Context) ([]entity.DBTablePost, error) {
	db := library.GetDB(ctx)
	return repository.GetDeletedPost_List(ctx, db)
}

// RestorePost は削除済みデータを復元します。
func RestorePost(ctx context.Context, id int64) error {
	db := library.GetDB(ctx)
	return repository.RestorePost(ctx, db, id)
}
