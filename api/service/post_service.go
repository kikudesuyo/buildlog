package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

func GetPostList(ctx context.Context) []entity.DBTablePost {
	return repository.GetPostList()
}
