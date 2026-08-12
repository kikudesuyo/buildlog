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

func withTrashTestDB(t *testing.T) sqlmock.Sqlmock {
	t.Helper()
	db, mock := testutil.NewTestDB(t)
	restore := library.SetDBForTest(db)
	t.Cleanup(restore)
	return mock
}

// TestListDeletedPostsReturnsPostsOrderedByDeletedAt は削除済み投稿を削除日時の降順で
// 取得できることを検証します。
func TestListDeletedPostsReturnsPostsOrderedByDeletedAt(t *testing.T) {
	mock := withTrashTestDB(t)
	deletedAt := time.Date(2026, 8, 12, 9, 0, 0, 0, time.UTC)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posts" WHERE deleted_at IS NOT NULL ORDER BY deleted_at DESC`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "type", "title", "deleted_at"}).
			AddRow(2, "diary", "newer", deletedAt).
			AddRow(1, "tech", "older", deletedAt.Add(-time.Hour)))

	posts, err := GetDeletedPostList(context.Background())
	if err != nil {
		t.Fatalf("ListDeletedPosts() error = %v", err)
	}
	if len(posts) != 2 || posts[0].ID != 2 || posts[1].ID != 1 {
		t.Fatalf("ListDeletedPosts() = %#v, want posts ordered by deletion time", posts)
	}
}

// TestListDeletedPostsReturnsEmptySliceWhenNoDeletedPostsExist は削除済み投稿がない場合に
// エラーなしの空スライスを返す境界値を検証します。
func TestListDeletedPostsReturnsEmptySliceWhenNoDeletedPostsExist(t *testing.T) {
	mock := withTrashTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posts" WHERE deleted_at IS NOT NULL ORDER BY deleted_at DESC`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "type", "title", "deleted_at"}))

	posts, err := GetDeletedPostList(context.Background())
	if err != nil {
		t.Fatalf("ListDeletedPosts() error = %v", err)
	}
	if posts == nil || len(posts) != 0 {
		t.Fatalf("ListDeletedPosts() = %#v, want non-nil empty slice", posts)
	}
}

// TestListDeletedPostsReturnsDatabaseError は削除済み投稿の取得に失敗した場合にDBエラーを
// 呼び出し元へ返すことを検証します。
func TestListDeletedPostsReturnsDatabaseError(t *testing.T) {
	mock := withTrashTestDB(t)
	databaseError := errors.New("database unavailable")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posts" WHERE deleted_at IS NOT NULL ORDER BY deleted_at DESC`)).
		WillReturnError(databaseError)

	posts, err := GetDeletedPostList(context.Background())
	if posts != nil || !errors.Is(err, databaseError) {
		t.Fatalf("ListDeletedPosts() = %#v, %v, want database error", posts, err)
	}
}

// TestRestorePostClearsDeletedAt は指定した投稿のdeleted_atをNULLに更新できることを検証します。
func TestRestorePostClearsDeletedAt(t *testing.T) {
	mock := withTrashTestDB(t)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "posts" SET "deleted_at"=$1,"updated_at"=$2 WHERE id = $3`)).
		WithArgs(nil, sqlmock.AnyArg(), int64(42)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	if err := RestorePost(context.Background(), 42); err != nil {
		t.Fatalf("RestorePost() error = %v", err)
	}
}

// TestRestorePostReturnsDatabaseError は復元対象の更新に失敗した場合にDBエラーを返すことを
// 検証します。
func TestRestorePostReturnsDatabaseError(t *testing.T) {
	mock := withTrashTestDB(t)
	databaseError := errors.New("restore failed")
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "posts" SET "deleted_at"=$1,"updated_at"=$2 WHERE id = $3`)).
		WithArgs(nil, sqlmock.AnyArg(), int64(42)).
		WillReturnError(databaseError)
	mock.ExpectRollback()

	if err := RestorePost(context.Background(), 42); !errors.Is(err, databaseError) {
		t.Fatalf("RestorePost() error = %v, want database error", err)
	}
}
