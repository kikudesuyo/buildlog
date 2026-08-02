package entity

import "time"

type DBTableGoalPeriod struct {
	ID         int64         `gorm:"column:id;primaryKey"`
	PeriodType string        `gorm:"column:period_type"`
	StartsAt   time.Time     `gorm:"column:starts_at"`
	EndsAt     time.Time     `gorm:"column:ends_at"`
	Goals      []DBTableGoal `gorm:"foreignKey:PeriodID"`
}

func (DBTableGoalPeriod) TableName() string { return "goal_periods" }

type DBTableGoal struct {
	ID            int64     `gorm:"column:id;primaryKey" json:"id"`
	PeriodID      int64     `gorm:"column:period_id" json:"period_id"`
	Title         string    `gorm:"column:title" json:"title"`
	TargetValue   int       `gorm:"column:target_value" json:"target_value"`
	ProgressValue int       `gorm:"column:progress_value" json:"progress_value"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (DBTableGoal) TableName() string { return "goals" }
