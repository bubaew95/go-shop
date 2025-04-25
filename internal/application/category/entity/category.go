package entity

import "time"

type Category struct {
	ID        int
	Name      string
	Alias     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
