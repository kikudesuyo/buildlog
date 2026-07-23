package entity

type CreateTechRequest struct {
	Title        string `json:"title"`
	Excerpt      string `json:"excerpt"`
	Category     string `json:"category"`
	ReadTime     string `json:"read_time"`
	Views        string `json:"views"`
	IsNewsletter bool   `json:"is_newsletter"`
}

type CreateTechResponse struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Excerpt      string `json:"excerpt"`
	Category     string `json:"category"`
	ReadTime     string `json:"read_time"`
	Views        string `json:"views"`
	IsNewsletter bool   `json:"is_newsletter"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type UpdateTechRequest struct {
	Title        string `json:"title"`
	Excerpt      string `json:"excerpt"`
	Category     string `json:"category"`
	ReadTime     string `json:"read_time"`
	Views        string `json:"views"`
	IsNewsletter bool   `json:"is_newsletter"`
}

type UpdateTechResponse struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Excerpt      string `json:"excerpt"`
	Category     string `json:"category"`
	ReadTime     string `json:"read_time"`
	Views        string `json:"views"`
	IsNewsletter bool   `json:"is_newsletter"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
