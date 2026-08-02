package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"github.com/kikudesuyo/buildlog/api/xerror"
	"gorm.io/gorm"
)

const monthlyGoalPeriod = "monthly"

func GetCurrentGoals(ctx context.Context) (*entity.GoalPeriodResponse, error) {
	startsAt := currentMonthStart()
	period, err := repository.GetGoalPeriod(ctx, database, monthlyGoalPeriod, startsAt)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return newGoalPeriodResponse(monthlyGoalPeriod, startsAt, startsAt.AddDate(0, 1, -1), nil), nil
	}
	if err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	return newGoalPeriodResponse(period.PeriodType, period.StartsAt, period.EndsAt, period.Goals), nil
}

func SaveCurrentGoals(ctx context.Context, req entity.SaveGoalsRequest) (*entity.GoalPeriodResponse, error) {
	if req.PeriodType != "" && req.PeriodType != monthlyGoalPeriod {
		return nil, xerror.ClientValidationErr(errors.New("unsupported goal period"))
	}
	if len(req.Goals) == 0 {
		return nil, xerror.ClientValidationErr(errors.New("at least one goal is required"))
	}

	startsAt := currentMonthStart()
	endsAt := startsAt.AddDate(0, 1, -1)
	goalList := make([]entity.DBTableGoal, 0, len(req.Goals))
	for _, goal := range req.Goals {
		title := strings.TrimSpace(goal.Title)
		if title == "" || goal.TargetValue <= 0 || goal.ProgressValue < 0 || goal.ProgressValue > goal.TargetValue {
			return nil, xerror.ClientValidationErr(errors.New("invalid goal values"))
		}
		goalList = append(goalList, entity.DBTableGoal{Title: title, TargetValue: goal.TargetValue, ProgressValue: goal.ProgressValue})
	}

	period, err := repository.GetGoalPeriod(ctx, database, monthlyGoalPeriod, startsAt)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		period = &entity.DBTableGoalPeriod{PeriodType: monthlyGoalPeriod, StartsAt: startsAt, EndsAt: endsAt}
	} else if err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	period.EndsAt = endsAt
	period.Goals = goalList
	if err := repository.SaveGoalPeriod(ctx, database, period); err != nil {
		return nil, xerror.UnknownServerErr(err)
	}
	return newGoalPeriodResponse(period.PeriodType, period.StartsAt, period.EndsAt, period.Goals), nil
}

func currentMonthStart() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
}

func newGoalPeriodResponse(periodType string, startsAt, endsAt time.Time, goalList []entity.DBTableGoal) *entity.GoalPeriodResponse {
	goals := make([]entity.GoalResponse, 0, len(goalList))
	for _, goal := range goalList {
		goals = append(goals, entity.GoalResponse{ID: goal.ID, Title: goal.Title, TargetValue: goal.TargetValue, ProgressValue: goal.ProgressValue})
	}
	return &entity.GoalPeriodResponse{PeriodType: periodType, StartsAt: startsAt.Format("2006-01-02"), EndsAt: endsAt.Format("2006-01-02"), Goals: goals}
}
