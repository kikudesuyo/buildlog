package service

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/testutil"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

func withCommentTestDB(t *testing.T) sqlmock.Sqlmock {
	t.Helper()
	db, mock := testutil.NewMockDB(t)
	restore := library.SetDBForTest(db)
	t.Cleanup(restore)
	return mock
}

func TestListCommentsByPostIDReturnsCommentsOrderedByCreation(t *testing.T) {
	mock := withCommentTestDB(t)
	createdAt := time.Date(2026, 8, 12, 9, 0, 0, 0, time.UTC)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE post_id = $1 ORDER BY created_at ASC,id ASC`)).
		WithArgs(int64(42)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "post_id", "content", "created_at", "updated_at"}).
			AddRow(1, 42, "first", createdAt, createdAt).
			AddRow(2, 42, "second", createdAt.Add(time.Minute), createdAt.Add(time.Minute)))

	comments, err := ListCommentsByPostID(context.Background(), 42)
	if err != nil {
		t.Fatalf("ListCommentsByPostID() error = %v", err)
	}
	if len(comments) != 2 || comments[0].Content != "first" || comments[1].Content != "second" {
		t.Fatalf("ListCommentsByPostID() = %#v, want two ordered comments", comments)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

func TestListCommentsByPostIDReturnsEmptySliceWhenNoCommentsExist(t *testing.T) {
	mock := withCommentTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE post_id = $1 ORDER BY created_at ASC,id ASC`)).
		WithArgs(int64(0)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "post_id", "content", "created_at", "updated_at"}))

	comments, err := ListCommentsByPostID(context.Background(), 0)
	if err != nil {
		t.Fatalf("ListCommentsByPostID() error = %v", err)
	}
	if comments == nil || len(comments) != 0 {
		t.Fatalf("ListCommentsByPostID() = %#v, want non-nil empty slice", comments)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

func TestListCommentsByPostIDWrapsRepositoryError(t *testing.T) {
	mock := withCommentTestDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "comments" WHERE post_id = $1 ORDER BY created_at ASC,id ASC`)).
		WithArgs(int64(7)).
		WillReturnError(errors.New("database unavailable"))

	comments, err := ListCommentsByPostID(context.Background(), 7)
	if comments != nil || err == nil || !xerror.IsCustomError(err) {
		t.Fatalf("ListCommentsByPostID() = %#v, %v, want custom error", comments, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

func TestCreateCommentPersistsContentAndReturnsComment(t *testing.T) {
	mock := withCommentTestDB(t)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "comments" ("post_id","content","created_at","updated_at") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WithArgs(int64(42), "hello", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(9))
	mock.ExpectCommit()

	comment, err := CreateComment(context.Background(), 42, entity.CreateCommentRequest{Content: "hello"})
	if err != nil {
		t.Fatalf("CreateComment() error = %v", err)
	}
	if comment == nil || comment.ID != 9 || comment.PostID != 42 || comment.Content != "hello" {
		t.Fatalf("CreateComment() = %#v, want persisted comment", comment)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

func TestCreateCommentWrapsRepositoryError(t *testing.T) {
	mock := withCommentTestDB(t)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "comments" ("post_id","content","created_at","updated_at") VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WithArgs(int64(42), "", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnError(errors.New("insert failed"))
	mock.ExpectRollback()

	comment, err := CreateComment(context.Background(), 42, entity.CreateCommentRequest{})
	if comment != nil || err == nil || !xerror.IsCustomError(err) {
		t.Fatalf("CreateComment() = %#v, %v, want custom error", comment, err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}
