package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
)

// ListPostHistory は削除されていない投稿履歴を取得します。
func ListPostHistory(ctx context.Context) ([]entity.HistoryItem, error) {
	return repository.ListPostHistory(ctx, library.GetDB(ctx))
}
