package entity

import "time"

type TechFeedItem struct {
	Key           string        `json:"key"`
	ID            int64         `json:"id"`
	Type          string        `json:"type"`
	Title         string        `json:"title"`
	Content       string        `json:"content"`
	Views         int64         `json:"views"`
	Status        string        `json:"status"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	LikesCount    int64         `json:"likes_count"`
	CommentsCount int64         `json:"comments_count"`
	HasLiked      bool          `json:"has_liked"`
	External      *ExternalPost `json:"external,omitempty"`
}

type ExternalPost struct {
	Provider     string `json:"provider"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnail_url"`
}
