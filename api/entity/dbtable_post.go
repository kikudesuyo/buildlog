package entity

import "time"

type DBTablePost struct {
	ID           int64     `gorm:"column:id;primaryKey" json:"id"`
	Type         string    `gorm:"column:type" json:"type"` // "diary" または "tech"
	Title        string    `gorm:"column:title" json:"title"`
	Content      string    `gorm:"column:content" json:"content"`
	Excerpt      string    `gorm:"column:excerpt" json:"excerpt"`
	Category     string    `gorm:"column:category" json:"category"`
	ReadTime     string    `gorm:"column:read_time" json:"read_time"`
	Views        string    `gorm:"column:views" json:"views"`
	IsNewsletter bool      `gorm:"column:is_newsletter" json:"is_newsletter"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (DBTablePost) TableName() string {
	return "posts"
}
