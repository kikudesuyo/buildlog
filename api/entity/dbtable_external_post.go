package entity

import "time"

type DBTableExternalPost struct {
	ID           int64     `gorm:"column:id;primaryKey" json:"id"`
	Provider     string    `gorm:"column:provider" json:"provider"`
	ExternalID   string    `gorm:"column:external_id" json:"external_id"`
	URL          string    `gorm:"column:url" json:"url"`
	Title        string    `gorm:"column:title" json:"title"`
	Excerpt      string    `gorm:"column:excerpt" json:"excerpt"`
	ThumbnailURL string    `gorm:"column:thumbnail_url" json:"thumbnail_url"`
	PublishedAt  time.Time `gorm:"column:published_at" json:"published_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
}

func (DBTableExternalPost) TableName() string { return "external_posts" }
