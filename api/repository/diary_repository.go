package repository

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// GetDiary_List は日記一覧を取得し、公開一覧ではページ単位で絞り込みます。
func GetDiary_List(ctx context.Context, db *gorm.DB, all bool, offset int, limit int, sortBy string, sortOrder string) ([]entity.DBTablePost, error) {
	diaryList := make([]entity.DBTablePost, 0)
	query := db.WithContext(ctx).Where("type = ?", "diary")
	if !all {
		query = query.Where("status = ?", "published")
	}
	if sortBy == "likes" {
		query = query.Order("COALESCE((SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id), 0) " + sortOrder).Order("id DESC")
	} else {
		query = query.Order("created_at " + sortOrder).Order("id " + sortOrder)
	}
	if !all {
		query = query.Offset(offset)
		if limit > 0 {
			query = query.Limit(limit)
		}
	}
	err := query.Find(&diaryList).Error
	return diaryList, err
}

// GetDiaryByID はデータを取得します。
func GetDiaryByID(ctx context.Context, db *gorm.DB, id int64) (*entity.DBTablePost, error) {
	var diary entity.DBTablePost
	err := db.WithContext(ctx).Where("type = ?", "diary").First(&diary, id).Error
	return &diary, err
}

// IncrementPostViews は投稿の閲覧数を 1 増やします。
func IncrementPostViews(ctx context.Context, db *gorm.DB, postType string, id int64) error {
	return db.WithContext(ctx).Model(&entity.DBTablePost{}).
		Where("type = ? AND id = ?", postType, id).
		UpdateColumn("views", gorm.Expr("COALESCE(views, 0) + 1")).Error
}

// CreateDiary はデータを作成します。
func CreateDiary(ctx context.Context, db *gorm.DB, diary *entity.DBTablePost) error {
	diary.Type = "diary"
	return db.WithContext(ctx).Create(diary).Error
}

// UpdateDiary はデータを更新します。
func UpdateDiary(ctx context.Context, db *gorm.DB, diary *entity.DBTablePost) error {
	diary.Type = "diary"
	return db.WithContext(ctx).Save(diary).Error
}

// DeleteDiary はデータを削除します。
func DeleteDiary(ctx context.Context, db *gorm.DB, id int64) error {
	return db.WithContext(ctx).Where("type = ?", "diary").Delete(&entity.DBTablePost{}, id).Error
}
