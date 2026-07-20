package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

func ListAppProjects(ctx context.Context) []entity.AppProject {
	return repository.ListAppProjects()
}
