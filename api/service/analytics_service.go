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
	itemList := make([]entity.AnalyticsArticleItem, len(data.Posts))
	for i, post := range data.Posts {
		views := post.Views
		totalViews += views
		switch post.Type {
		case "diary":
			diaryCount++
		case "tech":
			techCount++
		}
		itemList[i] = entity.AnalyticsArticleItem{ID: post.ID, Type: post.Type, Title: post.Title, Views: views, Likes: data.LikeCounts[post.ID]}
	}
	viewRankingList := append([]entity.AnalyticsArticleItem(nil), itemList...)
	sort.Slice(viewRankingList, func(i, j int) bool {
		if viewRankingList[i].Views == viewRankingList[j].Views {
			return viewRankingList[i].ID > viewRankingList[j].ID
		}
		return viewRankingList[i].Views > viewRankingList[j].Views
	})
	if len(viewRankingList) > 5 {
		viewRankingList = viewRankingList[:5]
	}
	likeRankingList := append([]entity.AnalyticsArticleItem(nil), itemList...)
	sort.Slice(likeRankingList, func(i, j int) bool {
		if likeRankingList[i].Likes == likeRankingList[j].Likes {
			return likeRankingList[i].ID > likeRankingList[j].ID
		}
		return likeRankingList[i].Likes > likeRankingList[j].Likes
	})
	if len(likeRankingList) > 5 {
		likeRankingList = likeRankingList[:5]
	}

	now := time.Now()
	monthList := make([]string, 12)
	countList := make(map[string]int64)
	for i := 11; i >= 0; i-- {
		month := now.AddDate(0, -i, 0).Format("2006-01")
		monthList[11-i] = month
		countList[month] = 0
	}
	for _, post := range data.Posts {
		month := post.CreatedAt.Format("2006-01")
		if _, ok := countList[month]; ok {
			countList[month]++
		}
	}
	activityList := make([]entity.MonthlyActivityItem, 12)
	for i, month := range monthList {
		activityList[i] = entity.MonthlyActivityItem{Month: month, Count: countList[month]}
	}
	return entity.AnalyticsResponse{TotalViews: totalViews, TotalLikes: data.TotalLikes, TotalPosts: int64(len(data.Posts)), DiaryCount: diaryCount, TechCount: techCount, TopViewsArticles: viewRankingList, TopLikesArticles: likeRankingList, MonthlyActivities: activityList}, nil
}
