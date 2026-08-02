package entity

import "time"

type DBTableComment struct {
	ID        int64             `gorm:"column:id;primaryKey" json:"id"`
	PostID    int64             `gorm:"column:post_id" json:"post_id"`
	ParentID  *int64            `gorm:"column:parent_id" json:"parent_id"`
	Content   string            `gorm:"column:content" json:"content"`
	CreatedAt time.Time         `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time         `gorm:"column:updated_at" json:"updated_at"`
	Replies   []*DBTableComment `gorm:"-" json:"replies,omitempty"`
}

func (DBTableComment) TableName() string {
	return "comments"
}
