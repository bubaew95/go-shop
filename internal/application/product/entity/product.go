package entity

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID             int       `json:"id" db:"id"`
	FirmID         int       `json:"firm_id" db:"firm_id"`
	UserID         int       `json:"user_id" db:"user_id"`
	Name           string    `json:"name" db:"name" validate:"required"`
	Anons          string    `json:"anons,omitempty" db:"anons"`
	Text           string    `json:"text,omitempty" db:"text"`
	Stock          int       `json:"stock,omitempty" db:"stock"`
	Price          float64   `json:"price" db:"price" validate:"required"`
	Discount       int       `json:"discount,omitempty" db:"discount"`
	SeoTitle       string    `json:"seo_title,omitempty" db:"seo_title"`
	SeoDescription string    `json:"seo_description,omitempty" db:"seo_description"`
	SeoKeywords    string    `json:"seo_keywords,omitempty" db:"seo_keywords"`
	Active         bool      `json:"active,omitempty" db:"active"`
	CreatedAt      time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

func (p *Product) Validate() (map[string][]string, bool) {
	errorsList := make(map[string][]string)

	err := validator.New().Struct(p)
	if err == nil {
		return nil, true
	}

	var errs validator.ValidationErrors
	if ok := errors.As(err, &errs); !ok {
		return nil, true
	}

	for _, err := range errs {
		field := strings.ToLower(err.Field())
		switch err.Tag() {
		case "required":
			errorsList[field] = append(errorsList[field], "is required")
		case "min":
			errorsList[field] = append(errorsList[field], fmt.Sprintf("length must be at least %s", err.Param()))
		case "gt":
			errorsList[field] = append(errorsList[field], "must be greater than zero")
		default:
			errorsList[field] = append(errorsList[field], fmt.Sprintf("failed on %s", err.Tag()))
		}
	}

	return errorsList, false
}
