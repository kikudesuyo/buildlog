package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

func GetTechFeed(ctx context.Context) entity.TechFeed {
	return repository.GetTechFeed()
}
