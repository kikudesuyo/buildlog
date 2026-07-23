package entity

type CreateDiaryRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateDiaryResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateDiaryRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateDiaryResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
