package v1

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/handler"
	"gorm.io/gorm"
)

type PostLikeCount struct {
	PostID int64 `gorm:"column:post_id"`
	Count  int64 `gorm:"column:count"`
}

func HandleGetAnalytics(db *gorm.DB) handler.ProcessFunc {
	return func(r *http.Request, requestData map[string]interface{}) (handler.Renderer, error) {
		var posts []entity.DBTablePost
		err := db.Where("deleted_at IS NULL").Find(&posts).Error
		if err != nil {
			return nil, err
		}

		// いいねの総件数を取得
		var totalLikes int64
		err = db.Model(&entity.DBTableLike{}).Count(&totalLikes).Error
		if err != nil {
			return nil, err
		}

		// 記事ごとのいいね数を集計
		var likeCounts []PostLikeCount
		err = db.Model(&entity.DBTableLike{}).Select("post_id, COUNT(*) as count").Group("post_id").Scan(&likeCounts).Error
		if err != nil {
			return nil, err
		}

		likesMap := make(map[int64]int64)
		for _, lc := range likeCounts {
			likesMap[lc.PostID] = lc.Count
		}

		var totalViews int64
		var diaryCount int64
		var techCount int64

		articleItems := make([]entity.AnalyticsArticleItem, len(posts))

		for i, post := range posts {
			viewsVal, _ := strconv.ParseInt(post.Views, 10, 64)
			likesVal := likesMap[post.ID]

			totalViews += viewsVal

			if post.Type == "diary" {
				diaryCount++
			} else if post.Type == "tech" {
				techCount++
			}

			articleItems[i] = entity.AnalyticsArticleItem{
				ID:    post.ID,
				Type:  post.Type,
				Title: post.Title,
				Views: viewsVal,
				Likes: likesVal,
			}
		}

		// 閲覧数ランキングの上位5件
		viewsRanking := make([]entity.AnalyticsArticleItem, len(articleItems))
		copy(viewsRanking, articleItems)
		sort.Slice(viewsRanking, func(i, j int) bool {
			if viewsRanking[i].Views == viewsRanking[j].Views {
				return viewsRanking[i].ID > viewsRanking[j].ID
			}
			return viewsRanking[i].Views > viewsRanking[j].Views
		})
		if len(viewsRanking) > 5 {
			viewsRanking = viewsRanking[:5]
		}

		// いいね数ランキングの上位5件
		likesRanking := make([]entity.AnalyticsArticleItem, len(articleItems))
		copy(likesRanking, articleItems)
		sort.Slice(likesRanking, func(i, j int) bool {
			if likesRanking[i].Likes == likesRanking[j].Likes {
				return likesRanking[i].ID > likesRanking[j].ID
			}
			return likesRanking[i].Likes > likesRanking[j].Likes
		})
		if len(likesRanking) > 5 {
			likesRanking = likesRanking[:5]
		}

		// 過去12ヶ月の月別投稿数集計
		now := time.Now()
		monthlyCounts := make(map[string]int64)
		monthsList := make([]string, 12)

		// 過去12ヶ月のキーを生成 (古い順に並べる)
		for i := 11; i >= 0; i-- {
			t := now.AddDate(0, -i, 0)
			monthStr := t.Format("2006-01")
			monthsList[11-i] = monthStr
			monthlyCounts[monthStr] = 0
		}

		for _, post := range posts {
			monthStr := post.CreatedAt.Format("2006-01")
			if _, exists := monthlyCounts[monthStr]; exists {
				monthlyCounts[monthStr]++
			}
		}

		monthlyActivities := make([]entity.MonthlyActivityItem, 12)
		for i, mStr := range monthsList {
			monthlyActivities[i] = entity.MonthlyActivityItem{
				Month: mStr,
				Count: monthlyCounts[mStr],
			}
		}

		resp := entity.AnalyticsResponse{
			TotalViews:        totalViews,
			TotalLikes:        totalLikes,
			TotalPosts:        int64(len(posts)),
			DiaryCount:        diaryCount,
			TechCount:         techCount,
			TopViewsArticles:  viewsRanking,
			TopLikesArticles:  likesRanking,
			MonthlyActivities: monthlyActivities,
		}

		return entity.NewObjectResponse(resp), nil
	}
}
