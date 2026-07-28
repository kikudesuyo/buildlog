package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

func ListDeletedPosts(ctx context.Context) ([]entity.DBTablePost, error) {
	return repository.ListDeletedPosts(ctx, DB)
}

func RestorePost(ctx context.Context, id int64) error {
	return repository.RestorePost(ctx, DB, id)
}
