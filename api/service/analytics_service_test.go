package service

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/testutil"
)

func withAnalyticsTestDB(t *testing.T) sqlmock.Sqlmock {
	t.Helper()
	db, mock := testutil.NewMockDB(t)
	restore := library.SetDBForTest(db)
	t.Cleanup(restore)
	return mock
}

func expectAnalyticsQueries(mock sqlmock.Sqlmock, posts *sqlmock.Rows, likes *sqlmock.Rows) {
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posts" WHERE deleted_at IS NULL`)).
		WillReturnRows(posts)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT post_id, COUNT(*) AS count FROM "likes" GROUP BY "post_id"`)).
		WillReturnRows(likes)
}

// TestGetAnalyticsAggregatesCountsAndLimitsRankings は、投稿・いいねの集計と
// 上位5件へのランキング制限を検証します。
func TestGetAnalyticsAggregatesCountsAndLimitsRankings(t *testing.T) {
	mock := withAnalyticsTestDB(t)
	now := time.Now()
	posts := sqlmock.NewRows([]string{"id", "type", "title", "views", "created_at", "updated_at", "deleted_at"})
	for i := int64(1); i <= 6; i++ {
		postType := "tech"
		if i%2 == 0 {
			postType = "diary"
		}
		posts.AddRow(i, postType, "post", i*10, now.AddDate(0, 0, -int(i)), now, nil)
	}
	likes := sqlmock.NewRows([]string{"post_id", "count"}).
		AddRow(1, 6).
		AddRow(2, 5).
		AddRow(3, 4).
		AddRow(4, 3).
		AddRow(5, 2).
		AddRow(6, 1)
	expectAnalyticsQueries(mock, posts, likes)

	result, err := GetAnalytics(context.Background())
	if err != nil {
		t.Fatalf("GetAnalytics() error = %v", err)
	}
	if result.TotalViews != 210 || result.TotalLikes != 21 || result.TotalPosts != 6 {
		t.Fatalf("summary = %#v, want views=210 likes=21 posts=6", result)
	}
	if result.DiaryCount != 3 || result.TechCount != 3 {
		t.Fatalf("type counts = diary:%d tech:%d, want 3 and 3", result.DiaryCount, result.TechCount)
	}
	if len(result.TopViewsArticles) != 5 || result.TopViewsArticles[0].ID != 6 {
		t.Fatalf("view ranking = %#v, want top five descending by views", result.TopViewsArticles)
	}
	if len(result.TopLikesArticles) != 5 || result.TopLikesArticles[0].ID != 1 {
		t.Fatalf("like ranking = %#v, want top five descending by likes", result.TopLikesArticles)
	}
	if len(result.MonthlyActivities) != 12 {
		t.Fatalf("monthly activities length = %d, want 12", len(result.MonthlyActivities))
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

// TestGetAnalyticsReturnsZeroValuesForEmptyData は、データがない場合に集計値が
// 0となり、月次活動だけが12か月分返る境界値を検証します。
func TestGetAnalyticsReturnsZeroValuesForEmptyData(t *testing.T) {
	mock := withAnalyticsTestDB(t)
	expectAnalyticsQueries(
		mock,
		sqlmock.NewRows([]string{"id", "type", "title", "views", "created_at", "updated_at", "deleted_at"}),
		sqlmock.NewRows([]string{"post_id", "count"}),
	)

	result, err := GetAnalytics(context.Background())
	if err != nil {
		t.Fatalf("GetAnalytics() error = %v", err)
	}
	if result.TotalViews != 0 || result.TotalLikes != 0 || result.TotalPosts != 0 || len(result.TopViewsArticles) != 0 || len(result.TopLikesArticles) != 0 {
		t.Fatalf("empty result = %#v, want zero summary and rankings", result)
	}
	if len(result.MonthlyActivities) != 12 {
		t.Fatalf("monthly activities length = %d, want 12", len(result.MonthlyActivities))
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

// TestGetAnalyticsReturnsRepositoryError は、投稿取得時のDBエラーが呼び出し元へ
// 返ることを検証します。
func TestGetAnalyticsReturnsRepositoryError(t *testing.T) {
	mock := withAnalyticsTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posts" WHERE deleted_at IS NULL`)).
		WillReturnError(errors.New("database unavailable"))

	result, err := GetAnalytics(context.Background())
	if err == nil {
		t.Fatal("GetAnalytics() error = nil, want repository error")
	}
	if result.TotalViews != 0 || result.TotalLikes != 0 || result.TotalPosts != 0 {
		t.Fatalf("result = %#v, want zero response on error", result)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}
