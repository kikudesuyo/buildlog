package entity

import (
	"time"

	"gorm.io/gorm"
)

type DBTableComment struct {
	ID        int64          `gorm:"column:id;primaryKey" json:"id"`
	PostID    int64          `gorm:"column:post_id" json:"post_id"`
	Content   string         `gorm:"column:content" json:"content"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at,omitempty"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName はエンティティに対応するデータベーステーブル名を返します。
func (DBTableComment) TableName() string {
	return "comments"
}
