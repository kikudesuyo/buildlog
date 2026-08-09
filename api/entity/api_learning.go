package entity

type CreateLearningRequest struct {
	PeriodStart string  `json:"period_start"`
	Content     string  `json:"content"`
	Level       *string `json:"level"`
}

type GenerateLearningRequest struct {
	PeriodStart string `json:"period_start"`
}

type LearningResponse struct {
	ID          int64  `json:"id"`
	PeriodType  string `json:"period_type"`
	PeriodStart string `json:"period_start"`
	PeriodEnd   string `json:"period_end"`
	Content     string `json:"content"`
	Level       string `json:"level,omitempty"`
	GeneratedBy string `json:"generated_by"`
}
