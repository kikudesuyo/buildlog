package testutil

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewMockDB creates a GORM database backed by sqlmock. It never connects to
// an external database and closes the underlying sql.DB at test cleanup.
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
		t.Fatalf("open GORM mock database: %v", err)
	}

	t.Cleanup(func() {
		_ = sqlDB.Close()
	})

	return db, mock
}
