package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
)

// GetPostHistoryList は削除されていない投稿履歴を取得します。
func GetPostHistoryList(ctx context.Context) ([]entity.HistoryItem, error) {
	return repository.GetPostHistoryList(ctx, library.GetDB(ctx))
}
