package entity

import "time"

type DBTablePost struct {
	ID        int64     `gorm:"column:id;primaryKey" json:"id"`
	Type      string    `gorm:"column:type" json:"type"` // "diary" または "tech"
	Title     string    `gorm:"column:title" json:"title"`
	Content   string    `gorm:"column:content" json:"content"`
	Category  string    `gorm:"column:category" json:"category"`
	Views     string    `gorm:"column:views" json:"views"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	LikesCount int64     `gorm:"-" json:"likes_count"`
	HasLiked   bool      `gorm:"-" json:"has_liked"`
}

func (DBTablePost) TableName() string {
	return "posts"
}
