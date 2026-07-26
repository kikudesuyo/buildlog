package entity

type UpdateProfileRequest struct {
	Name         string             `json:"name"`
	Subtitle     string             `json:"subtitle"`
	Title        string             `json:"title"`
	AvatarURL    string             `json:"avatar_url"`
	Quote        string             `json:"quote"`
	Bio          []string           `json:"bio"`
	Highlights   []ProfileHighlight `json:"highlights"`
	Award        string             `json:"award"`
	Expertise    []string           `json:"expertise"`
	ContactEmail string             `json:"contact_email"`
	FinalQuote   string             `json:"final_quote"`
}
