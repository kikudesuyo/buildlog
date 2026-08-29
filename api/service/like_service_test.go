package service

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/testutil"
)

func withLikeTestDB(t *testing.T) sqlmock.Sqlmock {
	t.Helper()
	db, mock := testutil.NewMockDB(t)
	t.Cleanup(library.SetDBForTest(db))
	return mock
}

// TestGetLikeStatusReturnsCountAndViewerState は、投稿にいいねがあり、
// 現在の閲覧者もいいね済みである通常系を検証します。
func TestGetLikeStatusReturnsCountAndViewerState(t *testing.T) {
	mock := withLikeTestDB(t)
	wantPostID := int64(42)
	wantIPAddress := "203.0.113.10"
	wantLikesCount := int64(3)
	wantHasLiked := true
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) AS likes_count, COALESCE(BOOL_OR(ip_address = $1), false) AS has_liked FROM "likes" WHERE post_id = $2`)).
		WithArgs(wantIPAddress, wantPostID).
		WillReturnRows(sqlmock.NewRows([]string{"likes_count", "has_liked"}).AddRow(wantLikesCount, wantHasLiked))

	status, err := GetLikeStatus(context.Background(), wantPostID, wantIPAddress)
	if err != nil {
		t.Fatalf("GetLikeStatus() error = %v", err)
	}
	if status.LikesCount != wantLikesCount || status.HasLiked != wantHasLiked {
		t.Fatalf("GetLikeStatus() = %#v, want likes_count=%d and has_liked=%t", status, wantLikesCount, wantHasLiked)
	}
}

// TestGetLikeStatusReturnsZeroForPostWithoutLikes は、投稿にいいねがなく、
// 閲覧者もいいねしていない境界値を検証します。
func TestGetLikeStatusReturnsZeroForPostWithoutLikes(t *testing.T) {
	mock := withLikeTestDB(t)
	wantPostID := int64(99)
	wantIPAddress := "198.51.100.20"
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) AS likes_count, COALESCE(BOOL_OR(ip_address = $1), false) AS has_liked FROM "likes" WHERE post_id = $2`)).
		WithArgs(wantIPAddress, wantPostID).
		WillReturnRows(sqlmock.NewRows([]string{"likes_count", "has_liked"}).AddRow(0, false))

	status, err := GetLikeStatus(context.Background(), wantPostID, wantIPAddress)
	if err != nil || status.LikesCount != 0 || status.HasLiked {
		t.Fatalf("GetLikeStatus() = %#v, %v, want zero count and not liked", status, err)
	}
}

// TestLikePostReturnsCreateError は、いいね作成時のエラーによって、
// 後続の状態取得が実行されず処理が終了することを検証します。
func TestLikePostReturnsCreateError(t *testing.T) {
	mock := withLikeTestDB(t)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "likes" ("post_id","ip_address","created_at") VALUES ($1,$2,$3) ON CONFLICT DO NOTHING`)).
		WithArgs(int64(42), "203.0.113.10", sqlmock.AnyArg()).
		WillReturnError(errors.New("insert failed"))
	mock.ExpectRollback()

	status, err := LikePost(context.Background(), 42, "203.0.113.10")
	if status != (LikeStatus{}) || err == nil {
		t.Fatalf("LikePost() = %#v, %v, want zero status and error", status, err)
	}
}

// TestUnlikePostReturnsDeleteError は、いいね削除時のエラーによって、
// 後続の状態取得が実行されず処理が終了することを検証します。
func TestUnlikePostReturnsDeleteError(t *testing.T) {
	mock := withLikeTestDB(t)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "likes" WHERE post_id = $1 AND ip_address = $2`)).
		WithArgs(int64(42), "203.0.113.10").
		WillReturnError(errors.New("delete failed"))
	mock.ExpectRollback()

	status, err := UnlikePost(context.Background(), 42, "203.0.113.10")
	if status != (LikeStatus{}) || err == nil {
		t.Fatalf("UnlikePost() = %#v, %v, want zero status and error", status, err)
	}
}
