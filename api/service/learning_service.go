package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/repository"
	"github.com/kikudesuyo/buildlog/api/xerror"
	"gorm.io/gorm"
)

const (
	DailyLearning   = "daily"
	WeeklyLearning  = "weekly"
	MonthlyLearning = "monthly"
	dateLayout      = "2006-01-02"
)

var validLearningLevels = map[string]bool{"learned": true, "understood": true, "applied": true, "explainable": true}

func FetchCurrentLearnings(ctx context.Context, periodType string, now time.Time) ([]entity.LearningResponse, error) {
	if !validPeriodType(periodType) {
		return nil, xerror.ClientValidationErr(errors.New("invalid learning period type"))
	}
	start, end := learningPeriod(periodType, now.UTC())
	items, err := repository.FetchLearnings(ctx, library.GetDB(ctx), periodType, start, end)
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

// CreateLearningSummary creates a period summary from the sources defined by the growth-log specification.
func CreateLearningSummary(ctx context.Context, periodType string, req entity.CreateLearningSummaryRequest) (*entity.LearningResponse, error) {
	if periodType != WeeklyLearning && periodType != MonthlyLearning {
		return nil, xerror.ClientValidationErr(errors.New("only weekly or monthly learning can be generated"))
	}
	requestedStart, err := time.Parse(dateLayout, req.PeriodStart)
	if err != nil {
		return nil, xerror.ClientValidationErr(errors.New("invalid period start"))
	}
	start, end := learningPeriod(periodType, requestedStart)
	db := library.GetDB(ctx)
	lines, err := learningSourceLines(ctx, db, periodType, start, end)
	if err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	if len(lines) == 0 {
		return nil, xerror.ClientValidationErr(errors.New("no source content found"))
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

func learningSourceLines(ctx context.Context, db *gorm.DB, periodType string, start, end time.Time) ([]string, error) {
	lines := make([]string, 0)
	if periodType == WeeklyLearning {
		posts, err := repository.FetchPublishedPostsForLearning(ctx, db, start, end)
		if err != nil {
			return nil, err
		}
		for _, post := range posts {
			lines = append(lines, formatLearningSource(post.Type, post.Title, post.Content))
		}
		externalPosts, err := repository.FetchExternalPostsForLearning(ctx, db, start, end)
		if err != nil {
			return nil, err
		}
		for _, post := range externalPosts {
			lines = append(lines, formatLearningSource(post.Provider, post.Title, post.Excerpt))
		}
		return lines, nil
	}

	weekly, err := repository.FetchLearnings(ctx, db, WeeklyLearning, start, end)
	if err != nil {
		return nil, err
	}
	for _, learning := range weekly {
		lines = append(lines, formatLearningSource("weekly", "", learning.Content))
	}
	goals, err := repository.GetGoalPeriod(ctx, db, monthlyGoalPeriod, start)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if goals != nil {
		for _, goal := range goals.Goals {
			progress := fmt.Sprintf("進捗 %d/%d", goal.ProgressValue, goal.TargetValue)
			lines = append(lines, formatLearningSource("goal", goal.Title, progress))
		}
	}
	return lines, nil
}

func formatLearningSource(source, title, content string) string {
	content = strings.TrimSpace(content)
	if title = strings.TrimSpace(title); title != "" {
		return "- [" + source + "] " + title + ": " + content
	}
	return "- [" + source + "] " + content
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
