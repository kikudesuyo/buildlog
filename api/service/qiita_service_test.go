package service

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/external"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/testutil"
)

func withQiitaTestDB(t *testing.T) sqlmock.Sqlmock {
	t.Helper()
	db, mock := testutil.NewMockDB(t)
	t.Cleanup(library.SetDBForTest(db))
	return mock
}

// TestSortTechFeedItems は、Techフィードを昇順・降順に並べ替え、同じ日時ではIDで
// 安定して並ぶことを検証します。
func TestSortTechFeedItems(t *testing.T) {
	items := []entity.TechFeedItem{
		{ID: 1, CreatedAt: time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC)},
		{ID: 2, CreatedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	sortTechFeedItems(items, "asc")
	if items[0].ID != 2 || items[1].ID != 1 {
		t.Fatalf("ascending order = [%d, %d], want [2, 1]", items[0].ID, items[1].ID)
	}

	sortTechFeedItems(items, "desc")
	if items[0].ID != 1 || items[1].ID != 2 {
		t.Fatalf("descending order = [%d, %d], want [1, 2]", items[0].ID, items[1].ID)
	}
}

// TestListTechFeedMapsExternalPostsAndPaginates は、外部記事をフィード項目へ変換し、
// offsetとlimitでページングする通常系を検証します。
func TestListTechFeedMapsExternalPostsAndPaginates(t *testing.T) {
	mock := withQiitaTestDB(t)
	now := time.Now().UTC()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "external_posts" ORDER BY published_at DESC,id DESC`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "provider", "url", "title", "excerpt", "thumbnail_url", "likes_count", "published_at", "updated_at", "created_at"}).
			AddRow(1, "qiita", "https://example.com/1", "one", "excerpt", "thumb", 4, now, now, now).
			AddRow(2, "qiita", "https://example.com/2", "two", "excerpt", "thumb", 2, now.Add(-time.Hour), now, now))

	items, err := ListTechFeed(context.Background(), library.GetDB(context.Background()), false, 1, 1, "desc", "")
	if err != nil || len(items) != 1 || items[0].ID != 2 || items[0].Key != "external:2" {
		t.Fatalf("ListTechFeed() = %#v, %v, want paginated external item", items, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

// TestListTechFeedReturnsRepositoryError は、外部記事取得時のDBエラーを検証します。
func TestListTechFeedReturnsRepositoryError(t *testing.T) {
	mock := withQiitaTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "external_posts" ORDER BY published_at ASC,id ASC`)).
		WillReturnError(errors.New("database unavailable"))

	items, err := ListTechFeed(context.Background(), library.GetDB(context.Background()), false, 0, 10, "asc", "")
	if items != nil || err == nil {
		t.Fatalf("ListTechFeed() = %#v, %v, want error", items, err)
	}
}

// TestExcerptForPrefersMetadataAndNormalizesMarkdown は、OGP説明を優先し、説明がない場合は
// 本文の空白を正規化して抜粋を作ることを検証します。
func TestExcerptForPrefersMetadataAndNormalizesMarkdown(t *testing.T) {
	item := external.QiitaItem{Body: " first\n\n second "}
	if got := excerptFor(item, external.OGPMetadata{Description: "metadata"}); got != "metadata" {
		t.Fatalf("excerptFor() = %q, want metadata", got)
	}
	if got := excerptFor(item, external.OGPMetadata{}); got != "first second" {
		t.Fatalf("excerptFor() = %q, want normalized body", got)
	}
}

// TestSyncQiitaArticlesReturnsContextErrorWithoutExternalCall は、キャンセル済みContextでは
// 外部同期を実行せずエラーを返すことを検証します。
func TestSyncQiitaArticlesReturnsContextErrorWithoutExternalCall(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	count, err := SyncQiitaArticles(ctx)
	if count != 0 || err == nil {
		t.Fatalf("SyncQiitaArticles() = %d, %v, want canceled-context error", count, err)
	}
}
