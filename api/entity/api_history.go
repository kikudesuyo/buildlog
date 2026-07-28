package entity

import "time"

type HistoryItem struct {
	ID        int64     `json:"id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}


