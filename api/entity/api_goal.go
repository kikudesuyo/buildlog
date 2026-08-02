package entity

type SaveGoalsRequest struct {
	PeriodType string            `json:"period_type"`
	Goals      []SaveGoalRequest `json:"goals"`
}

type SaveGoalRequest struct {
	Title         string `json:"title"`
	TargetValue   int    `json:"target_value"`
	ProgressValue int    `json:"progress_value"`
}

type GoalPeriodResponse struct {
	PeriodType string         `json:"period_type"`
	StartsAt   string         `json:"starts_at"`
	EndsAt     string         `json:"ends_at"`
	Goals      []GoalResponse `json:"goals"`
}

type GoalResponse struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	TargetValue   int    `json:"target_value"`
	ProgressValue int    `json:"progress_value"`
}
