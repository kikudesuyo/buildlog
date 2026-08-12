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

func withProfileTestDB(t *testing.T) sqlmock.Sqlmock {
	t.Helper()
	db, mock := testutil.NewMockDB(t)
	restore := library.SetDBForTest(db)
	t.Cleanup(restore)
	return mock
}

func profileQuery(mock sqlmock.Sqlmock) *sqlmock.ExpectedQuery {
	return mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "profiles" WHERE "profiles"."id" = $1 ORDER BY "profiles"."id" LIMIT $2`)).WithArgs(1, 1)
}

// TestGetProfileDecodesJSONFieldsAndAssetURL は、プロフィールのJSON列をレスポンス型へ
// 変換する通常系を検証します。
func TestGetProfileDecodesJSONFieldsAndAssetURL(t *testing.T) {
	mock := withProfileTestDB(t)
	profileQuery(mock).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "title", "quote", "bio", "highlights", "award", "expertise", "contact_email", "final_quote"}).
		AddRow(1, "Ada", "Engineer", "quote", `["one","two"]`, `[{"title":"Award","period":"2026","description":"desc"}]`, "award", `["Go"]`, "ada@example.com", "final"))

	result, err := GetProfile(context.Background())
	if err != nil {
		t.Fatalf("GetProfile() error = %v", err)
	}
	if result.Name != "Ada" || len(result.Bio) != 2 || len(result.Highlights) != 1 || result.Expertise[0] != "Go" {
		t.Fatalf("GetProfile() = %#v, want decoded profile", result)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

// TestGetProfileReturnsJSONDecodeError は、プロフィールのJSON列が不正な場合を検証します。
func TestGetProfileReturnsJSONDecodeError(t *testing.T) {
	mock := withProfileTestDB(t)
	profileQuery(mock).WillReturnRows(sqlmock.NewRows([]string{"id", "bio"}).AddRow(1, "not-json"))

	result, err := GetProfile(context.Background())
	if result != nil || err == nil {
		t.Fatalf("GetProfile() = %#v, %v, want decode error", result, err)
	}
}

// TestUpdateProfileReturnsLookupErrorBeforeSaving は、更新前のプロフィール取得に失敗した場合に
// 保存処理へ進まないことを検証します。
func TestUpdateProfileReturnsLookupErrorBeforeSaving(t *testing.T) {
	mock := withProfileTestDB(t)
	profileQuery(mock).WillReturnError(errors.New("profile unavailable"))

	result, err := UpdateProfile(context.Background(), entity.UpdateProfileRequest{Name: "new"})
	if result != nil || err == nil {
		t.Fatalf("UpdateProfile() = %#v, %v, want lookup error", result, err)
	}
}
