package entity

import "time"

type DBTablePost struct {
	ID        int64     `gorm:"column:id"`
	Title     string    `gorm:"column:title"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (DBTablePost) TableName() string {
	return "posts"
}
