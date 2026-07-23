package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func ListDiaryEntries(ctx context.Context, db *gorm.DB) ([]entity.DiaryEntry, error) {
	return repository.ListDiaryEntries(ctx, db)
}
