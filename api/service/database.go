package service

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type databaseContextKey struct{}

// WithDatabase はリクエストコンテキストへ DB 依存を注入します。
func WithDatabase(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, databaseContextKey{}, db)
}

func databaseFromContext(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(databaseContextKey{}).(*gorm.DB)
	if !ok || db == nil {
		panic(errors.New("database is not configured in context"))
	}
	return db
}
