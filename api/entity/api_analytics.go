package entity

type AnalyticsResponse struct {
	TotalViews        int64                  `json:"total_views"`
	TotalLikes        int64                  `json:"total_likes"`
	TotalPosts        int64                  `json:"total_posts"`
	DiaryCount        int64                  `json:"diary_count"`
	TechCount         int64                  `json:"tech_count"`
	TopViewsArticles  []AnalyticsArticleItem `json:"top_views_articles"`
	TopLikesArticles  []AnalyticsArticleItem `json:"top_likes_articles"`
	MonthlyActivities []MonthlyActivityItem  `json:"monthly_activities"`
}

type AnalyticsArticleItem struct {
	ID    int64  `json:"id"`
	Type  string `json:"type"`
	Title string `json:"title"`
	Views int64  `json:"views"`
	Likes int64  `json:"likes"`
}

type MonthlyActivityItem struct {
	Month string `json:"month"` // YYYY-MM
	Count int64  `json:"count"`
}
