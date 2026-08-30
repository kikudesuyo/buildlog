package v1

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/testutil"
)

// TestHandleGetPostHistoryAllowsUnauthenticatedRequests は公開プロフィールが利用する
// 投稿履歴GET APIを、認証情報なしで取得できることを検証します。
func TestHandleGetPostHistoryAllowsUnauthenticatedRequests(t *testing.T) {
	db, mock := testutil.NewMockDB(t)
	t.Cleanup(library.SetDBForTest(db))
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, type, title, created_at, '' AS url FROM posts WHERE deleted_at IS NULL UNION ALL SELECT id, 'tech' AS type, title, published_at AS created_at, url FROM external_posts ORDER BY created_at DESC`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "type", "title", "created_at", "url"}))

	request := httptest.NewRequest(http.MethodGet, "/api/v1/posts/history", nil)
	response, err := HandleGetPostHistory(request, nil)
	if err != nil {
		t.Fatalf("HandleGetPostHistory() error = %v", err)
	}

	recorder := httptest.NewRecorder()
	response.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", recorder.Code, http.StatusOK)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}
