package service

import (
	"context"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func ListDiaries(ctx context.Context, db *gorm.DB, tag string) ([]entity.DBTablePost, error) {
	return repository.ListDiaries(ctx, db, tag)
}

func GetDiaryByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTablePost, error) {
	return repository.GetDiaryByID(ctx, db, id)
}

func CreateDiary(ctx context.Context, db *gorm.DB, req entity.CreateDiaryRequest) (entity.CreateDiaryResponse, error) {
	diary := entity.DBTablePost{
		Title:   req.Title,
		Content: req.Content,
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := repository.CreateDiary(ctx, tx, &diary); err != nil {
			return err
		}
		if err := savePostTags(tx, &diary, req.Tags); err != nil {
			return err
		}
		return tx.Preload("Tags").First(&diary, diary.ID).Error
	})
	if err != nil {
		return entity.CreateDiaryResponse{}, err
	}

	return entity.CreateDiaryResponse{
		ID:        diary.ID,
		Title:     diary.Title,
		Content:   diary.Content,
		Tags:      mapTagsToStrings(diary.Tags),
		CreatedAt: diary.CreatedAt.Format(time.RFC3339),
		UpdatedAt: diary.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func UpdateDiary(ctx context.Context, db *gorm.DB, id int64, req entity.UpdateDiaryRequest) (entity.UpdateDiaryResponse, error) {
	var diary *entity.DBTablePost
	err := db.Transaction(func(tx *gorm.DB) error {
		var err error
		diary, err = repository.GetDiaryByID(ctx, tx, id)
		if err != nil {
			return err
		}

		diary.Title = req.Title
		diary.Content = req.Content

		if err := repository.UpdateDiary(ctx, tx, diary); err != nil {
			return err
		}
		if err := savePostTags(tx, diary, req.Tags); err != nil {
			return err
		}
		return tx.Preload("Tags").First(diary, diary.ID).Error
	})
	if err != nil {
		return entity.UpdateDiaryResponse{}, err
	}

	return entity.UpdateDiaryResponse{
		ID:        diary.ID,
		Title:     diary.Title,
		Content:   diary.Content,
		Tags:      mapTagsToStrings(diary.Tags),
		CreatedAt: diary.CreatedAt.Format(time.RFC3339),
		UpdatedAt: diary.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func DeleteDiary(ctx context.Context, db *gorm.DB, id int64) error {
	return repository.DeleteDiary(ctx, db, id)
}

func savePostTags(tx *gorm.DB, post *entity.DBTablePost, tagNames []string) error {
	tags := make([]entity.DBTableTag, 0, len(tagNames))
	for _, name := range tagNames {
		if name == "" {
			continue
		}
		var t entity.DBTableTag
		if err := tx.Where("name = ?", name).FirstOrCreate(&t, entity.DBTableTag{Name: name}).Error; err != nil {
			return err
		}
		tags = append(tags, t)
	}
	return tx.Model(post).Association("Tags").Replace(tags)
}

func mapTagsToStrings(tags []entity.DBTableTag) []string {
	names := make([]string, len(tags))
	for i, t := range tags {
		names[i] = t.Name
	}
	return names
}
