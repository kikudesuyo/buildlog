package entity

import "time"

type DBTableLearning struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`
	PeriodType  string    `gorm:"column:period_type" json:"period_type"`
	PeriodStart time.Time `gorm:"column:period_start" json:"period_start"`
	PeriodEnd   time.Time `gorm:"column:period_end" json:"period_end"`
	Content     string    `gorm:"column:content" json:"content"`
	Level       *string   `gorm:"column:level" json:"level"`
	GeneratedBy string    `gorm:"column:generated_by" json:"generated_by"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (DBTableLearning) TableName() string { return "learnings" }
