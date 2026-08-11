package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
)

// GetTech_List は外部記事の一覧を取得します。
func GetTech_List(ctx context.Context, all bool, offset, limit int, order, ipAddress string) ([]entity.TechFeedItem, error) {
	return GetTechFeed_List(ctx, library.GetDB(ctx), all, offset, limit, order, ipAddress)
}
