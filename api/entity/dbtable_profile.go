package entity

import "time"

type DBTableProfile struct {
	ID           int64     `gorm:"column:id;primaryKey"`
	Name         string    `gorm:"column:name"`
	Title        string    `gorm:"column:title"`
	AvatarURL    string    `gorm:"column:avatar_url"`
	Quote        string    `gorm:"column:quote"`
	Bio          string    `gorm:"column:bio"`
	Highlights   string    `gorm:"column:highlights"`
	Award        string    `gorm:"column:award"`
	Expertise    string    `gorm:"column:expertise"`
	ContactEmail string    `gorm:"column:contact_email"`
	FinalQuote   string    `gorm:"column:final_quote"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

// TableName はエンティティに対応するデータベーステーブル名を返します。
func (DBTableProfile) TableName() string {
	return "profiles"
}
