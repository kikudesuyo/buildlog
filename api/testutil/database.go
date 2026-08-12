package testutil

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewTestDB creates a GORM database backed by a sqlmock connection.
// The connection and outstanding expectations are checked when the test ends.
func NewTestDB(t testing.TB) (*gorm.DB, sqlmock.Sqlmock) {
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
		mock.ExpectClose()
		if err := closeTestDB(db); err != nil {
			t.Errorf("close test database: %v", err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unmet SQL mock expectations: %v", err)
		}
	})

	return db, mock
}

func closeTestDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
