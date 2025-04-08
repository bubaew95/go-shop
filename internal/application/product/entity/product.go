package entity

type Product struct {
	ID     int     `json:"id" db:"id"`
	Name   string  `json:"name" db:"name"`
	Price  float64 `json:"price" db:"price"`
	Anons  *string `json:"anons,omitempty" db:"anons"`
	TEXT   string  `json:"text" db:"text"`
	Sale   float32 `json:"sale,omitempty" db:"sale"`
	Active bool    `json:"active,omitempty" db:"active"`
}
