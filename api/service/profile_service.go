package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

func GetProfileData(ctx context.Context) entity.ProfileData {
	return repository.GetProfileData()
}
