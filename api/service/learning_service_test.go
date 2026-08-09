package service

import (
	"testing"
	"time"
)

func TestLearningPeriod(t *testing.T) {
	now := time.Date(2026, time.August, 12, 14, 30, 0, 0, time.FixedZone("JST", 9*60*60))
	tests := []struct {
		periodType string
		start      string
		end        string
	}{
		{DailyLearning, "2026-08-12", "2026-08-12"},
		{WeeklyLearning, "2026-08-10", "2026-08-16"},
		{MonthlyLearning, "2026-08-01", "2026-08-31"},
	}
	for _, testCase := range tests {
		start, end := learningPeriod(testCase.periodType, now)
		if start.Format(dateLayout) != testCase.start || end.Format(dateLayout) != testCase.end {
			t.Errorf("learningPeriod(%q) = %s, %s; want %s, %s", testCase.periodType, start.Format(dateLayout), end.Format(dateLayout), testCase.start, testCase.end)
		}
	}
}
