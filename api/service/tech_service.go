package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
)

// ListTechs は外部記事の一覧を取得します。
func ListTechs(ctx context.Context, all bool, offset, limit int, ipAddress string) ([]entity.TechFeedItem, error) {
	return ListTechFeed(ctx, library.GetDB(ctx), all, offset, limit, ipAddress)
}
