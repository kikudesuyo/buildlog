package entity

import "time"

type DBTableLike struct {
	ID        int64     `gorm:"column:id;primaryKey" json:"id"`
	PostID    int64     `gorm:"column:post_id" json:"post_id"`
	IPAddress string    `gorm:"column:ip_address" json:"ip_address"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (DBTableLike) TableName() string {
	return "likes"
}
