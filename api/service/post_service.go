package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func GetPostList(ctx context.Context, db *gorm.DB) ([]entity.DBTablePost, error) {
	return repository.GetPostList(ctx, db)
}
