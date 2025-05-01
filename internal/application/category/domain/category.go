package domain

import (
	"context"

	"github.com/bubaew95/go_shop/internal/application/category/entity"
)

type CategoryRepository interface {
	Create(ctx context.Context, category entity.Category) (entity.Category, error)
	GetCategories(ctx context.Context, offset int, limit int) ([]entity.CategoryResponse, error)
}

type CategoryService interface {
	Create(ctx context.Context, category entity.Category) (entity.Category, error)
	GetCategories(ctx context.Context, offset int, limit int) ([]entity.CategoryResponse, error)
}
