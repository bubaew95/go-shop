package entity

import "time"

type Category struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type Images struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Base  bool   `json:"is_base"`
}

type ProductResponse struct {
	ID             int        `json:"id"`
	FirmID         *int       `json:"firm_id"`
	Name           string     `json:"name"`
	Anons          string     `json:"anons,omitempty"`
	Text           string     `json:"text,omitempty"`
	Stock          int        `json:"stock,omitempty"`
	Price          float64    `json:"price"`
	Discount       int        `json:"discount,omitempty"`
	SeoTitle       string     `json:"seo_title,omitempty"`
	SeoDescription string     `json:"seo_description,omitempty"`
	SeoKeywords    string     `json:"seo_keywords,omitempty"`
	CreatedAt      time.Time  `json:"created_at,omitempty"`
	Category       []Category `json:"category"`
	Images         []Images   `json:"images"`
}
