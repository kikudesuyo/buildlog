package testutil

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// TestNewTestDBUsesSQLMockConnection は、テスト用DBがSQL mock接続を使い、外部DBへ接続せず
// SQL期待値を検証できることを確認します。
func TestNewTestDBUsesSQLMockConnection(t *testing.T) {
	db, mock := NewTestDB(t)
	mock.ExpectExec(regexp.QuoteMeta("SELECT 1")).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := db.Exec("SELECT 1").Error; err != nil {
		t.Fatalf("execute test query: %v", err)
	}
}
