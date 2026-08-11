package testutil

import (
	"regexp"
	"testing"
)

func TestNewMockDBUsesGORMAndSQLMock(t *testing.T) {
	db, mock := NewMockDB(t)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT 1`)).WillReturnRows(mock.NewRows([]string{"?column?"}).AddRow(1))

	var value int
	if err := db.Raw("SELECT 1").Scan(&value).Error; err != nil {
		t.Fatalf("mock query failed: %v", err)
	}
	if value != 1 {
		t.Fatalf("value = %d, want 1", value)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet SQL expectations: %v", err)
	}
}
