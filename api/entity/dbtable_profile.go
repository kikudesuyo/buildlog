package entity

import "time"

type ProfileHighlight struct {
	Title       string `json:"title"`
	Period      string `json:"period"`
	Description string `json:"description"`
}

type DBTableProfile struct {
	ID           int64              `gorm:"column:id;primaryKey" json:"id"`
	Name         string             `gorm:"column:name" json:"name"`
	Subtitle     string             `gorm:"column:subtitle" json:"subtitle"`
	Title        string             `gorm:"column:title" json:"title"`
	AvatarURL    string             `gorm:"column:avatar_url" json:"avatar_url"`
	Quote        string             `gorm:"column:quote" json:"quote"`
	Bio          []string           `gorm:"column:bio;serializer:json" json:"bio"`
	Highlights   []ProfileHighlight `gorm:"column:highlights;serializer:json" json:"highlights"`
	Award        string             `gorm:"column:award" json:"award"`
	Expertise    []string           `gorm:"column:expertise;serializer:json" json:"expertise"`
	ContactEmail string             `gorm:"column:contact_email" json:"contact_email"`
	FinalQuote   string             `gorm:"column:final_quote" json:"final_quote"`
	CreatedAt    time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time          `gorm:"column:updated_at" json:"updated_at"`
}

func (DBTableProfile) TableName() string {
	return "profiles"
}
