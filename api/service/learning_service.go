package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
	"github.com/kikudesuyo/buildlog/api/xerror"
)

const (
	DailyLearning   = "daily"
	WeeklyLearning  = "weekly"
	MonthlyLearning = "monthly"
	dateLayout      = "2006-01-02"
)

var validLearningLevels = map[string]bool{"learned": true, "understood": true, "applied": true, "explainable": true}

func GetCurrentLearnings(ctx context.Context, periodType string, now time.Time) ([]entity.LearningResponse, error) {
	if !validPeriodType(periodType) {
		return nil, xerror.ClientValidationErr(errors.New("invalid learning period type"))
	}
	start, end := learningPeriod(periodType, now.UTC())
	items, err := repository.ListLearnings(ctx, library.GetDB(ctx), periodType, start, end)
	if err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	responses := make([]entity.LearningResponse, 0, len(items))
	for _, item := range items {
		responses = append(responses, learningResponse(item))
	}
	return responses, nil
}

func CreateDailyLearning(ctx context.Context, req entity.CreateLearningRequest) (*entity.LearningResponse, error) {
	start, err := time.Parse(dateLayout, req.PeriodStart)
	content := strings.TrimSpace(req.Content)
	if err != nil || content == "" || req.Level == nil || !validLearningLevels[*req.Level] {
		return nil, xerror.ClientValidationErr(errors.New("invalid daily learning"))
	}
	learning := entity.DBTableLearning{
		PeriodType: DailyLearning, PeriodStart: start, PeriodEnd: start,
		Content: content, Level: req.Level, GeneratedBy: "user",
	}
	if err := repository.CreateLearning(ctx, library.GetDB(ctx), &learning); err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	response := learningResponse(learning)
	return &response, nil
}

// GenerateLearning creates a period summary from the available lower-granularity records.
// The generated text is intentionally limited to source content and does not infer new achievements.
func GenerateLearning(ctx context.Context, periodType string, req entity.GenerateLearningRequest) (*entity.LearningResponse, error) {
	if periodType != WeeklyLearning && periodType != MonthlyLearning {
		return nil, xerror.ClientValidationErr(errors.New("only weekly or monthly learning can be generated"))
	}
	requestedStart, err := time.Parse(dateLayout, req.PeriodStart)
	if err != nil {
		return nil, xerror.ClientValidationErr(errors.New("invalid period start"))
	}
	start, end := learningPeriod(periodType, requestedStart)
	sourceType := DailyLearning
	if periodType == MonthlyLearning {
		sourceType = WeeklyLearning
	}
	sources, err := repository.ListLearnings(ctx, library.GetDB(ctx), sourceType, start, end)
	if err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	if len(sources) == 0 {
		return nil, xerror.ClientValidationErr(errors.New("no source learnings found"))
	}
	lines := make([]string, 0, len(sources))
	for _, source := range sources {
		lines = append(lines, "- "+strings.TrimSpace(source.Content))
	}
	prefix := "今週取り組んだこと:"
	if periodType == MonthlyLearning {
		prefix = "今月の成長と理解:"
	}
	learning := entity.DBTableLearning{
		PeriodType: periodType, PeriodStart: start, PeriodEnd: end,
		Content: prefix + "\n" + strings.Join(lines, "\n"), GeneratedBy: "ai",
	}
	if err := repository.CreateLearning(ctx, library.GetDB(ctx), &learning); err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	response := learningResponse(learning)
	return &response, nil
}

func learningPeriod(periodType string, now time.Time) (time.Time, time.Time) {
	day := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	switch periodType {
	case DailyLearning:
		return day, day
	case WeeklyLearning:
		mondayOffset := (int(day.Weekday()) + 6) % 7
		start := day.AddDate(0, 0, -mondayOffset)
		return start, start.AddDate(0, 0, 6)
	default:
		start := time.Date(day.Year(), day.Month(), 1, 0, 0, 0, 0, time.UTC)
		return start, start.AddDate(0, 1, -1)
	}
}

func validPeriodType(value string) bool {
	return value == DailyLearning || value == WeeklyLearning || value == MonthlyLearning
}

func learningResponse(item entity.DBTableLearning) entity.LearningResponse {
	level := ""
	if item.Level != nil {
		level = *item.Level
	}
	return entity.LearningResponse{
		ID: item.ID, PeriodType: item.PeriodType,
		PeriodStart: item.PeriodStart.Format(dateLayout), PeriodEnd: item.PeriodEnd.Format(dateLayout),
		Content: item.Content, Level: level, GeneratedBy: item.GeneratedBy,
	}
}
