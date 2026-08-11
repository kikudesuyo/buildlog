package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
)

// GetPostHistory_List は削除されていない投稿履歴を取得します。
func GetPostHistory_List(ctx context.Context) ([]entity.HistoryItem, error) {
	return repository.GetPostHistory_List(ctx, library.GetDB(ctx))
}
