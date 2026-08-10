package entity

import "time"

type DBTablePostSource struct {
	ID         int64     `gorm:"column:id;primaryKey" json:"id"`
	PostID     int64     `gorm:"column:post_id" json:"post_id"`
	Provider   string    `gorm:"column:provider" json:"provider"`
	ExternalID string    `gorm:"column:external_id" json:"external_id"`
	URL        string    `gorm:"column:url" json:"url"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (DBTablePostSource) TableName() string { return "post_sources" }
