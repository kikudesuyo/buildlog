package entity

type AppResponse struct {
	ID          int64    `json:"id"`
	Slug        string   `json:"slug"`
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	Icon        string   `json:"icon"`
	IconURL     string   `json:"icon_url"`
	DemoURL     string   `json:"demo_url"`
	CodeURL     string   `json:"code_url"`
}

type CreateAppRequest struct {
	Slug        string   `json:"slug"`
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	Icon        string   `json:"icon"`
	IconURL     string   `json:"icon_url"`
	DemoURL     string   `json:"demo_url"`
	CodeURL     string   `json:"code_url"`
}

type UpdateAppRequest struct {
	Slug        string   `json:"slug"`
	Name        string   `json:"name"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	Icon        string   `json:"icon"`
	IconURL     string   `json:"icon_url"`
	DemoURL     string   `json:"demo_url"`
	CodeURL     string   `json:"code_url"`
}
