package service

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/testutil"
)

func withDiaryTestDB(t *testing.T) sqlmock.Sqlmock {
	t.Helper()
	db, mock := testutil.NewMockDB(t)
	restore := library.SetDBForTest(db)
	t.Cleanup(restore)
	return mock
}

func TestListDiariesReturnsEmptySliceForNoPublishedPosts(t *testing.T) {
	mock := withDiaryTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posts" WHERE type = $1 AND status = $2 AND "posts"."deleted_at" IS NULL ORDER BY created_at DESC,id DESC LIMIT $3`)).
		WithArgs("diary", "published", 10).
		WillReturnRows(sqlmock.NewRows([]string{"id", "type", "title", "content", "status"}))

	diaries, err := ListDiaries(context.Background(), false, 0, 10, "created_at", "DESC", "127.0.0.1")
	if err != nil || diaries == nil || len(diaries) != 0 {
		t.Fatalf("ListDiaries() = %#v, %v, want empty slice", diaries, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

func TestListDiariesReturnsRepositoryError(t *testing.T) {
	mock := withDiaryTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posts" WHERE type = $1 AND status = $2 AND "posts"."deleted_at" IS NULL ORDER BY created_at DESC,id DESC LIMIT $3`)).
		WithArgs("diary", "published", 10).
		WillReturnError(errors.New("list failed"))

	diaries, err := ListDiaries(context.Background(), false, 0, 10, "created_at", "DESC", "")
	if diaries != nil || err == nil {
		t.Fatalf("ListDiaries() = %#v, %v, want error", diaries, err)
	}
}

func TestGetDiaryByIDReturnsRepositoryError(t *testing.T) {
	mock := withDiaryTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posts" WHERE type = $1 AND "posts"."id" = $2 AND "posts"."deleted_at" IS NULL ORDER BY "posts"."id" LIMIT $3`)).
		WithArgs("diary", int64(8), 1).
		WillReturnError(errors.New("not found"))

	diary, err := GetDiaryByID(context.Background(), 8, "")
	if diary != nil || err == nil {
		t.Fatalf("GetDiaryByID() = %#v, %v, want repository error", diary, err)
	}
}

func TestIncrementDiaryViewsUpdatesDiaryOnly(t *testing.T) {
	mock := withDiaryTestDB(t)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "posts" SET "views"=COALESCE(views, 0) + 1 WHERE (type = $1 AND id = $2) AND "posts"."deleted_at" IS NULL`)).
		WithArgs("diary", int64(8)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	if err := IncrementDiaryViews(context.Background(), 8); err != nil {
		t.Fatalf("IncrementDiaryViews() error = %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

func TestCreateDiaryUsesDraftAsDefaultStatus(t *testing.T) {
	mock := withDiaryTestDB(t)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "posts" ("type","title","content","category","views","status","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id"`)).
		WithArgs("diary", "title", "body", "", int64(0), "draft", sqlmock.AnyArg(), sqlmock.AnyArg(), nil).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(12))
	mock.ExpectCommit()

	result, err := CreateDiary(context.Background(), entity.CreateDiaryRequest{Title: "title", Content: "body"})
	if err != nil || result.ID != 12 || result.Status != "draft" {
		t.Fatalf("CreateDiary() = %#v, %v, want draft diary", result, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

func TestUpdateDiaryReturnsLookupError(t *testing.T) {
	mock := withDiaryTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "posts" WHERE type = $1 AND "posts"."id" = $2 AND "posts"."deleted_at" IS NULL ORDER BY "posts"."id" LIMIT $3`)).
		WithArgs("diary", int64(3), 1).
		WillReturnError(errors.New("lookup failed"))

	result, err := UpdateDiary(context.Background(), 3, entity.UpdateDiaryRequest{Title: "new"})
	if result.ID != 0 || err == nil {
		t.Fatalf("UpdateDiary() = %#v, %v, want lookup error", result, err)
	}
}

func TestDeleteDiaryDeletesByIDAndType(t *testing.T) {
	mock := withDiaryTestDB(t)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "posts" SET "deleted_at"=$1 WHERE type = $2 AND "posts"."id" = $3 AND "posts"."deleted_at" IS NULL`)).
		WithArgs(sqlmock.AnyArg(), "diary", int64(4)).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	if err := DeleteDiary(context.Background(), 4); err != nil {
		t.Fatalf("DeleteDiary() error = %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}
