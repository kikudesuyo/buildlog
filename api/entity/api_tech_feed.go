package entity

import "time"

type TechFeedItem struct {
	Key       string       `json:"key"`
	ID        int64        `json:"id"`
	Title     string       `json:"title"`
	Excerpt   string       `json:"excerpt"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	External  ExternalPost `json:"external"`
}

type ExternalPost struct {
	Provider     string `json:"provider"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnail_url"`
}
