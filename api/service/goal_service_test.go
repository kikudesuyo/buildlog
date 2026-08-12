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
	"github.com/kikudesuyo/buildlog/api/xerror"
	"gorm.io/gorm"
)

func withGoalTestDB(t *testing.T) sqlmock.Sqlmock {
	t.Helper()
	db, mock := testutil.NewMockDB(t)
	restore := library.SetDBForTest(db)
	t.Cleanup(restore)
	return mock
}

func expectGoalPeriodLookup(mock sqlmock.Sqlmock) *sqlmock.ExpectedQuery {
	return mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "goal_periods" WHERE period_type = $1 AND starts_at = $2 ORDER BY "goal_periods"."id" LIMIT $3`)).
		WithArgs(monthlyGoalPeriod, sqlmock.AnyArg(), 1)
}

func TestGetCurrentGoalsReturnsEmptyCurrentMonthWhenNoPeriodExists(t *testing.T) {
	mock := withGoalTestDB(t)
	expectGoalPeriodLookup(mock).WillReturnError(gorm.ErrRecordNotFound)

	result, err := GetCurrentGoals(context.Background())
	if err != nil {
		t.Fatalf("GetCurrentGoals() error = %v", err)
	}
	if result == nil || result.PeriodType != monthlyGoalPeriod || len(result.Goals) != 0 {
		t.Fatalf("GetCurrentGoals() = %#v, want empty monthly period", result)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("SQL expectations: %v", err)
	}
}

func TestGetCurrentGoalsWrapsRepositoryError(t *testing.T) {
	mock := withGoalTestDB(t)
	expectGoalPeriodLookup(mock).WillReturnError(errors.New("database unavailable"))

	result, err := GetCurrentGoals(context.Background())
	if result != nil || err == nil || !xerror.IsCustomError(err) {
		t.Fatalf("GetCurrentGoals() = %#v, %v, want custom error", result, err)
	}
}

func TestSaveCurrentGoalsRejectsUnsupportedPeriod(t *testing.T) {
	withGoalTestDB(t)
	result, err := SaveCurrentGoals(context.Background(), entity.SaveGoalsRequest{PeriodType: "weekly", Goals: []entity.SaveGoalRequest{{Title: "goal", TargetValue: 1}}})
	if result != nil || err == nil || !xerror.IsCustomError(err) {
		t.Fatalf("SaveCurrentGoals() = %#v, %v, want validation error", result, err)
	}
}

func TestSaveCurrentGoalsRejectsInvalidGoalValues(t *testing.T) {
	withGoalTestDB(t)
	result, err := SaveCurrentGoals(context.Background(), entity.SaveGoalsRequest{Goals: []entity.SaveGoalRequest{{Title: " ", TargetValue: 0, ProgressValue: 1}}})
	if result != nil || err == nil || !xerror.IsCustomError(err) {
		t.Fatalf("SaveCurrentGoals() = %#v, %v, want validation error", result, err)
	}
}
