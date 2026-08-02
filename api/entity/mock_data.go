package entity

type DiaryEntry struct {
	ID                 string `json:"id"`
	Title              string `json:"title"`
	Excerpt            string `json:"excerpt"`
	Category           string `json:"category"`
	Date               string `json:"date"`
	Image              string `json:"image,omitempty"`
	ImageAlt           string `json:"image_alt,omitempty"`
	CategoryColorClass string `json:"category_color_class,omitempty"`
}

type TechArticle struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Excerpt  string `json:"excerpt"`
	Category string `json:"category"`
	Date     string `json:"date"`
	Views    string `json:"views,omitempty"`
}

type FeaturedTechArticle struct {
	Title    string `json:"title"`
	Excerpt  string `json:"excerpt"`
	Category string `json:"category"`
	Date     string `json:"date"`
}

type TechFeed struct {
	FeaturedArticle FeaturedTechArticle `json:"featured_article"`
	Articles        []TechArticle       `json:"articles"`
}
