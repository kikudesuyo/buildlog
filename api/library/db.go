package library

import (
	"context"
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// SetDBForTest replaces the process-local database used by services and
// returns a function that restores the previous value. It is intended for
// tests that use a sqlmock-backed GORM connection.
func SetDBForTest(testDB *gorm.DB) func() {
	previous := db
	db = testDB
	return func() {
		db = previous
	}
}

func InitDB() error {
	dsn := Env("DATABASE_URL")
	if dsn == "" {
		return errors.New("DATABASE_URL is not set")
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("get sql.DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	db = gormDB

	return nil
}

// Env はProductionでは既存の環境変数を、StagingではSTAGING_ prefix付きの
// 環境変数を返します。ローカル開発は既存の環境変数を利用します。
func Env(name string) string {
	if os.Getenv("IS_PRODUCTION") == "false" {
		return os.Getenv("STAGING_" + name)
	}
	return os.Getenv(name)
}

func GetDB(ctx context.Context) *gorm.DB {
	if ctx == nil {
		ctx = context.Background()
	}

	return db.WithContext(ctx)
}
