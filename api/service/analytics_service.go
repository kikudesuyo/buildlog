package service

import (
	"context"
	"sort"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
)

type PostLikeCount struct {
	PostID int64 `gorm:"column:post_id"`
	Count  int64 `gorm:"column:count"`
}

// GetAnalytics はデータを取得します。
func GetAnalytics(ctx context.Context) (entity.AnalyticsResponse, error) {
	db := library.GetDB(ctx)
	data, err := repository.GetAnalyticsData(ctx, db)
	if err != nil {
		return entity.AnalyticsResponse{}, err
	}

	var totalViews, diaryCount, techCount int64
	item_List := make([]entity.AnalyticsArticleItem, len(data.Posts))
	for i, post := range data.Posts {
		views := post.Views
		totalViews += views
		switch post.Type {
		case "diary":
			diaryCount++
		case "tech":
			techCount++
		}
		item_List[i] = entity.AnalyticsArticleItem{ID: post.ID, Type: post.Type, Title: post.Title, Views: views, Likes: data.LikeCounts[post.ID]}
	}
	viewRanking_List := append([]entity.AnalyticsArticleItem(nil), item_List...)
	sort.Slice(viewRanking_List, func(i, j int) bool {
		if viewRanking_List[i].Views == viewRanking_List[j].Views {
			return viewRanking_List[i].ID > viewRanking_List[j].ID
		}
		return viewRanking_List[i].Views > viewRanking_List[j].Views
	})
	if len(viewRanking_List) > 5 {
		viewRanking_List = viewRanking_List[:5]
	}
	likeRanking_List := append([]entity.AnalyticsArticleItem(nil), item_List...)
	sort.Slice(likeRanking_List, func(i, j int) bool {
		if likeRanking_List[i].Likes == likeRanking_List[j].Likes {
			return likeRanking_List[i].ID > likeRanking_List[j].ID
		}
		return likeRanking_List[i].Likes > likeRanking_List[j].Likes
	})
	if len(likeRanking_List) > 5 {
		likeRanking_List = likeRanking_List[:5]
	}

	now := time.Now()
	month_List := make([]string, 12)
	count_List := make(map[string]int64)
	for i := 11; i >= 0; i-- {
		month := now.AddDate(0, -i, 0).Format("2006-01")
		month_List[11-i] = month
		count_List[month] = 0
	}
	for _, post := range data.Posts {
		month := post.CreatedAt.Format("2006-01")
		if _, ok := count_List[month]; ok {
			count_List[month]++
		}
	}
	activity_List := make([]entity.MonthlyActivityItem, 12)
	for i, month := range month_List {
		activity_List[i] = entity.MonthlyActivityItem{Month: month, Count: count_List[month]}
	}
	return entity.AnalyticsResponse{TotalViews: totalViews, TotalLikes: data.TotalLikes, TotalPosts: int64(len(data.Posts)), DiaryCount: diaryCount, TechCount: techCount, TopViewsArticles: viewRanking_List, TopLikesArticles: likeRanking_List, MonthlyActivities: activity_List}, nil
}
