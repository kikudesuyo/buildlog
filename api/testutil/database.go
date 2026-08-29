package testutil

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewMockDB はsqlmockを使用したGORMのテスト用DBを生成します。
// 外部DBへ接続せず、テスト終了時に内部のsql.DBを解放します。
func NewMockDB(t testing.TB) (*gorm.DB, sqlmock.Sqlmock) {
	t.Helper()

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock database: %v", err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:       sqlDB,
		DriverName: "postgres",
	}), &gorm.Config{})
	if err != nil {
		_ = sqlDB.Close()
		t.Fatalf("open GORM database: %v", err)
	}

	t.Cleanup(func() {
		_ = sqlDB.Close()
	})

	return db, mock
}

// NewTestDB はsqlmock接続を使用したGORMのテスト用DBを生成します。
func NewTestDB(t testing.TB) (*gorm.DB, sqlmock.Sqlmock) {
	return NewMockDB(t)
}
