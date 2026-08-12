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

func withHistoryTestDB(t *testing.T) sqlmock.Sqlmock {
	t.Helper()
	db, mock := testutil.NewMockDB(t)
	t.Cleanup(library.SetDBForTest(db))
	return mock
}

// TestListPostHistoryReturnsNonDeletedPostsInDescendingOrder は、通常系として
// アプリ投稿と外部Qiita投稿を含む履歴が新しい順に返り、削除済み投稿が除外されることを検証します。
func TestListPostHistoryReturnsNonDeletedPostsInDescendingOrder(t *testing.T) {
	mock := withHistoryTestDB(t)
	newestCreatedAt := time.Date(2026, 8, 12, 10, 0, 0, 0, time.UTC)
	olderCreatedAt := newestCreatedAt.Add(-time.Hour)
	wantItemCount := 2
	wantNewestID := int64(2)
	wantOlderID := int64(1)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, type, title, created_at, '' AS url FROM posts WHERE deleted_at IS NULL UNION ALL SELECT id, 'tech' AS type, title, published_at AS created_at, url FROM external_posts ORDER BY created_at DESC`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "type", "title", "created_at", "url"}).
			AddRow(wantNewestID, "tech", "new", newestCreatedAt, "https://example.com/new").
			AddRow(wantOlderID, "diary", "old", olderCreatedAt, ""))

	items, err := ListPostHistory(context.Background())
	if err != nil {
		t.Fatalf("ListPostHistory() error = %v", err)
	}
	if len(items) != wantItemCount || items[0].ID != wantNewestID || items[0].Type != "tech" || items[1].ID != wantOlderID {
		t.Fatalf("ListPostHistory() = %#v, want %d items ordered newest-first with IDs [%d, %d]", items, wantItemCount, wantNewestID, wantOlderID)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

// TestListPostHistoryReturnsNonNilEmptySliceWhenNoPostsExist は、アプリ投稿と
// 外部投稿がどちらも存在しない境界値で、空の非nilスライスが返ることを検証します。
func TestListPostHistoryReturnsNonNilEmptySliceWhenNoPostsExist(t *testing.T) {
	mock := withHistoryTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, type, title, created_at, '' AS url FROM posts WHERE deleted_at IS NULL UNION ALL SELECT id, 'tech' AS type, title, published_at AS created_at, url FROM external_posts ORDER BY created_at DESC`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "type", "title", "created_at", "url"}))

	items, err := ListPostHistory(context.Background())
	wantItemCount := 0
	if err != nil || items == nil || len(items) != wantItemCount {
		t.Fatalf("ListPostHistory() = %#v, %v, want a non-nil slice with %d items", items, err, wantItemCount)
	}
}

// TestListPostHistoryReturnsRepositoryError は、データベース障害時に
// 初期化済みの空スライスとエラーが呼び出し元へ返ることを検証します。
func TestListPostHistoryReturnsRepositoryError(t *testing.T) {
	mock := withHistoryTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, type, title, created_at, '' AS url FROM posts WHERE deleted_at IS NULL UNION ALL SELECT id, 'tech' AS type, title, published_at AS created_at, url FROM external_posts ORDER BY created_at DESC`)).
		WillReturnError(errors.New("database unavailable"))

	items, err := ListPostHistory(context.Background())
	wantItemCount := 0
	if items == nil || len(items) != wantItemCount || err == nil {
		t.Fatalf("ListPostHistory() = %#v, %v, want a non-nil slice with %d items and an error", items, err, wantItemCount)
	}
}
