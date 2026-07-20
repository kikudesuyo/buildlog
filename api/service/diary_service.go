package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

func ListDiaryEntries(ctx context.Context) []entity.DiaryEntry {
	return repository.ListDiaryEntries()
}
