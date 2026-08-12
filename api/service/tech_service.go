package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
)

// GetTechList は外部記事の一覧を取得します。
func GetTechList(ctx context.Context, all bool, offset, limit int, order, ipAddress string) ([]entity.TechFeedItem, error) {
	return GetTechFeedList(ctx, library.GetDB(ctx), all, offset, limit, order, ipAddress)
}
