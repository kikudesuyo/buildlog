package entity

type CreateTechRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Views   int64  `json:"views"`
	Status  string `json:"status"`
}

type CreateTechResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Views     int64  `json:"views"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateTechRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Views   int64  `json:"views"`
	Status  string `json:"status"`
}

type UpdateTechResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Views     int64  `json:"views"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
