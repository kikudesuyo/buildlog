package repository

import (
	"context"
	"fmt"

	"github.com/kikudesuyo/buildlog/api/entity"
	"gorm.io/gorm"
)

// AnalyticsData は分析画面で使用するデータを表します。
type AnalyticsData struct {
	Posts      []entity.DBTablePost
	LikeCounts map[int64]int64
	TotalLikes int64
}

// GetAnalyticsData は分析に必要なデータをrepository層で取得します。
func GetAnalyticsData(ctx context.Context, db *gorm.DB) (AnalyticsData, error) {
	data := AnalyticsData{
		Posts:      make([]entity.DBTablePost, 0),
		LikeCounts: make(map[int64]int64),
	}
	if err := db.WithContext(ctx).Where("deleted_at IS NULL").Find(&data.Posts).Error; err != nil {
		return AnalyticsData{}, fmt.Errorf("list analytics posts: %w", err)
	}

	type count struct {
		PostID int64
		Count  int64
	}
	var countList []count
	if err := db.WithContext(ctx).
		Model(&entity.DBTableLike{}).
		Select("post_id, COUNT(*) AS count").
		Group("post_id").
		Find(&countList).Error; err != nil {
		return AnalyticsData{}, fmt.Errorf("list analytics like counts: %w", err)
	}
	for _, item := range countList {
		data.LikeCounts[item.PostID] = item.Count
		data.TotalLikes += item.Count
	}
	return data, nil
}
