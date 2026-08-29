package service

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/testutil"
)

var appColumns = []string{"id", "slug", "name", "category", "tags", "description", "icon", "icon_url", "demo_url", "code_url", "created_at", "updated_at"}

func useMockDB(t *testing.T) sqlmock.Sqlmock {
	t.Helper()
	db, mock := testutil.NewMockDB(t)
	restore := library.SetDBForTest(db)
	t.Cleanup(restore)
	return mock
}

func TestGetAppListReturnsMappedApps(t *testing.T) {
	mock := useMockDB(t)
	mock.ExpectQuery(`SELECT .* FROM "apps" ORDER BY id ASC`).WillReturnRows(
		sqlmock.NewRows(appColumns).AddRow(42, "sample-app", "Sample App", "tool", `["go"]`, "Description", "terminal", "icons/app.svg", "https://demo.example.test", "https://github.com/example/app", time.Now(), time.Now()),
	)

	got, err := GetAppList(context.Background())
	if err != nil {
		t.Fatalf("ListApps returned error: %v", err)
	}
	if len(got) != 1 || got[0].ID != 42 || got[0].Tags[0] != "go" {
		t.Fatalf("ListApps = %+v", got)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet SQL expectations: %v", err)
	}
}

func TestGetAppListReturnsRepositoryError(t *testing.T) {
	mock := useMockDB(t)
	mock.ExpectQuery(`SELECT .* FROM "apps" ORDER BY id ASC`).WillReturnError(sqlmock.ErrCancelled)

	if _, err := GetAppList(context.Background()); err == nil {
		t.Fatal("ListApps returned nil error")
	}
}

func TestGetAppByIDReturnsMappedApp(t *testing.T) {
	mock := useMockDB(t)
	mock.ExpectQuery(`SELECT .* FROM "apps" WHERE "apps"\."id" = \$1`).WithArgs(int64(42), 1).WillReturnRows(
		sqlmock.NewRows(appColumns).AddRow(42, "sample-app", "Sample App", "tool", `[]`, "Description", "terminal", "icons/app.svg", "", "", time.Now(), time.Now()),
	)

	got, err := GetAppByID(context.Background(), 42)
	if err != nil || got.ID != 42 {
		t.Fatalf("GetAppByID = %+v, %v", got, err)
	}
}

func TestCreateAppPersistsTagsAndReturnsApp(t *testing.T) {
	mock := useMockDB(t)
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "apps"`).WillReturnRows(
		sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(42, time.Now(), time.Now()),
	)
	mock.ExpectCommit()

	got, err := CreateApp(context.Background(), entity.CreateAppRequest{Slug: "sample-app", Name: "Sample App", Tags: []string{"go", "svelte"}})
	if err != nil {
		t.Fatalf("CreateApp returned error: %v", err)
	}
	if got.ID != 42 || len(got.Tags) != 2 {
		t.Fatalf("CreateApp = %+v", got)
	}
}

func TestUpdateAppUpdatesExistingApp(t *testing.T) {
	mock := useMockDB(t)
	mock.ExpectQuery(`SELECT .* FROM "apps" WHERE "apps"\."id" = \$1`).WithArgs(int64(42), 1).WillReturnRows(
		sqlmock.NewRows(appColumns).AddRow(42, "old", "Old", "tool", `[]`, "Old", "terminal", "", "", "", time.Now(), time.Now()),
	)
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "apps" SET`).WillReturnResult(sqlmock.NewResult(42, 1))
	mock.ExpectCommit()

	got, err := UpdateApp(context.Background(), 42, entity.UpdateAppRequest{Slug: "new", Name: "New", Tags: []string{"go"}})
	if err != nil || got.Slug != "new" || got.Name != "New" {
		t.Fatalf("UpdateApp = %+v, %v", got, err)
	}
}

func TestDeleteAppDeletesByID(t *testing.T) {
	mock := useMockDB(t)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "apps" WHERE "apps"."id" = $1`)).WithArgs(int64(42)).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	if err := DeleteApp(context.Background(), 42); err != nil {
		t.Fatalf("DeleteApp returned error: %v", err)
	}
}
