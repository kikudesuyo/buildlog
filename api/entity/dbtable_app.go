package entity

import "time"

type DBTableApp struct {
	ID          int64     `gorm:"column:id;primaryKey" json:"id"`
	Slug        string    `gorm:"column:slug;unique" json:"slug"`
	Name        string    `gorm:"column:name" json:"name"`
	Category    string    `gorm:"column:category" json:"category"`
	Tags        []string  `gorm:"column:tags;serializer:json" json:"tags"`
	Description string    `gorm:"column:description" json:"description"`
	Icon        string    `gorm:"column:icon" json:"icon"`
	IconURL     string    `gorm:"column:icon_url" json:"icon_url"`
	DemoURL     string    `gorm:"column:demo_url" json:"demo_url"`
	CodeURL     string    `gorm:"column:code_url" json:"code_url"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (DBTableApp) TableName() string {
	return "apps"
}
