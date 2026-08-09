package entity

type ProfileHighlight struct {
	Title       string `json:"title"`
	Period      string `json:"period"`
	Description string `json:"description"`
}

type ProfileResponse struct {
	Name         string             `json:"name"`
	Subtitle     string             `json:"subtitle"`
	Title        string             `json:"title"`
	Quote        string             `json:"quote"`
	Bio          []string           `json:"bio"`
	Highlights   []ProfileHighlight `json:"highlights"`
	Award        string             `json:"award"`
	Expertise    []string           `json:"expertise"`
	ContactEmail string             `json:"contact_email"`
	FinalQuote   string             `json:"final_quote"`
}

type UpdateProfileRequest struct {
	Name         string             `json:"name"`
	Subtitle     string             `json:"subtitle"`
	Title        string             `json:"title"`
	Quote        string             `json:"quote"`
	Bio          []string           `json:"bio"`
	Highlights   []ProfileHighlight `json:"highlights"`
	Award        string             `json:"award"`
	Expertise    []string           `json:"expertise"`
	ContactEmail string             `json:"contact_email"`
	FinalQuote   string             `json:"final_quote"`
}
