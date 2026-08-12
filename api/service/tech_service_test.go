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

func withTechTestDB(t *testing.T) sqlmock.Sqlmock {
	t.Helper()
	db, mock := testutil.NewMockDB(t)
	t.Cleanup(library.SetDBForTest(db))
	return mock
}

// TestListTechsReturnsExternalArticlesInRequestedOrder は、外部記事が
// 指定した並び順でTechFeedItemへ変換される通常系を検証します。
func TestListTechsReturnsExternalArticlesInRequestedOrder(t *testing.T) {
	mock := withTechTestDB(t)
	wantItemCount := 2
	wantFirstID := int64(2)
	wantFirstKey := "external:2"
	now := time.Date(2026, 8, 12, 10, 0, 0, 0, time.UTC)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "external_posts" ORDER BY published_at DESC,id DESC`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "provider", "url", "title", "excerpt", "thumbnail_url", "likes_count", "published_at", "updated_at", "created_at"}).
			AddRow(2, "qiita", "https://example.com/2", "new", "summary", "thumb", 4, now, now, now).
			AddRow(1, "qiita", "https://example.com/1", "old", "summary", "thumb", 1, now.Add(-time.Hour), now, now))

	items, err := ListTechs(context.Background(), false, 0, 10, "desc", "")
	if err != nil {
		t.Fatalf("ListTechs() error = %v", err)
	}
	if len(items) != wantItemCount || items[0].ID != wantFirstID || items[0].Key != wantFirstKey || items[0].External == nil {
		t.Fatalf("ListTechs() = %#v, want %d items with first ID=%d and key=%q", items, wantItemCount, wantFirstID, wantFirstKey)
	}
}

// TestListTechsReturnsEmptySliceWhenOffsetExceedsResults は、offsetが結果件数を
// 超える境界値で、エラーではなく空スライスが返ることを検証します。
func TestListTechsReturnsEmptySliceWhenOffsetExceedsResults(t *testing.T) {
	mock := withTechTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "external_posts" ORDER BY published_at ASC,id ASC`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "provider", "url", "title", "excerpt", "thumbnail_url", "likes_count", "published_at", "updated_at", "created_at"}))

	items, err := ListTechs(context.Background(), false, 1, 10, "asc", "")
	if err != nil || items == nil || len(items) != 0 {
		t.Fatalf("ListTechs() = %#v, %v, want non-nil empty slice when offset exceeds results", items, err)
	}
}

// TestListTechsReturnsRepositoryError は、外部記事取得時のDBエラーが
// 呼び出し元へ返ることを検証します。
func TestListTechsReturnsRepositoryError(t *testing.T) {
	mock := withTechTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "external_posts" ORDER BY published_at DESC,id DESC`)).
		WillReturnError(errors.New("database unavailable"))

	items, err := ListTechs(context.Background(), false, 0, 10, "desc", "")
	if items != nil || err == nil {
		t.Fatalf("ListTechs() = %#v, %v, want nil items and error", items, err)
	}
}
