package service

import (
	"context"
	"sort"
	"strconv"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
)

type PostLikeCount struct {
	PostID int64 `gorm:"column:post_id"`
	Count  int64 `gorm:"column:count"`
}

// GetAnalytics はデータを取得します。
func GetAnalytics(ctx context.Context) (entity.AnalyticsResponse, error) {
	var posts []entity.DBTablePost
	if err := databaseFromContext(ctx).WithContext(ctx).Where("deleted_at IS NULL").Find(&posts).Error; err != nil {
		return entity.AnalyticsResponse{}, err
	}
	var totalLikes int64
	if err := databaseFromContext(ctx).WithContext(ctx).Model(&entity.DBTableLike{}).Count(&totalLikes).Error; err != nil {
		return entity.AnalyticsResponse{}, err
	}
	var likeCounts []PostLikeCount
	if err := databaseFromContext(ctx).WithContext(ctx).Model(&entity.DBTableLike{}).Select("post_id, COUNT(*) as count").Group("post_id").Scan(&likeCounts).Error; err != nil {
		return entity.AnalyticsResponse{}, err
	}
	likesMap := make(map[int64]int64)
	for _, likeCount := range likeCounts {
		likesMap[likeCount.PostID] = likeCount.Count
	}

	var totalViews, diaryCount, techCount int64
	items := make([]entity.AnalyticsArticleItem, len(posts))
	for i, post := range posts {
		views, _ := strconv.ParseInt(post.Views, 10, 64)
		totalViews += views
		if post.Type == "diary" {
			diaryCount++
		} else if post.Type == "tech" {
			techCount++
		}
		items[i] = entity.AnalyticsArticleItem{ID: post.ID, Type: post.Type, Title: post.Title, Views: views, Likes: likesMap[post.ID]}
	}
	viewsRanking := append([]entity.AnalyticsArticleItem(nil), items...)
	sort.Slice(viewsRanking, func(i, j int) bool {
		if viewsRanking[i].Views == viewsRanking[j].Views {
			return viewsRanking[i].ID > viewsRanking[j].ID
		}
		return viewsRanking[i].Views > viewsRanking[j].Views
	})
	if len(viewsRanking) > 5 {
		viewsRanking = viewsRanking[:5]
	}
	likesRanking := append([]entity.AnalyticsArticleItem(nil), items...)
	sort.Slice(likesRanking, func(i, j int) bool {
		if likesRanking[i].Likes == likesRanking[j].Likes {
			return likesRanking[i].ID > likesRanking[j].ID
		}
		return likesRanking[i].Likes > likesRanking[j].Likes
	})
	if len(likesRanking) > 5 {
		likesRanking = likesRanking[:5]
	}

	now := time.Now()
	months := make([]string, 12)
	counts := make(map[string]int64)
	for i := 11; i >= 0; i-- {
		month := now.AddDate(0, -i, 0).Format("2006-01")
		months[11-i] = month
		counts[month] = 0
	}
	for _, post := range posts {
		month := post.CreatedAt.Format("2006-01")
		if _, ok := counts[month]; ok {
			counts[month]++
		}
	}
	activities := make([]entity.MonthlyActivityItem, 12)
	for i, month := range months {
		activities[i] = entity.MonthlyActivityItem{Month: month, Count: counts[month]}
	}
	return entity.AnalyticsResponse{TotalViews: totalViews, TotalLikes: totalLikes, TotalPosts: int64(len(posts)), DiaryCount: diaryCount, TechCount: techCount, TopViewsArticles: viewsRanking, TopLikesArticles: likesRanking, MonthlyActivities: activities}, nil
}
