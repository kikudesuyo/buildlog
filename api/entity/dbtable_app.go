package entity

import "time"

type DBTableApp struct {
	ID          int64     `gorm:"column:id;primaryKey"`
	Slug        string    `gorm:"column:slug;unique"`
	Name        string    `gorm:"column:name"`
	Category    string    `gorm:"column:category"`
	Tags        string    `gorm:"column:tags"`
	Description string    `gorm:"column:description"`
	Icon        string    `gorm:"column:icon"`
	IconURL     string    `gorm:"column:icon_url"`
	DemoURL     string    `gorm:"column:demo_url"`
	CodeURL     string    `gorm:"column:code_url"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

// TableName はエンティティに対応するデータベーステーブル名を返します。
func (DBTableApp) TableName() string {
	return "apps"
}
