package ports

import (
	"context"
	entity "github.com/bubaew95/go_shop/internal/core/entity/repository"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	GetProducts(ctx context.Context, offset int, limit int) ([]entity.Product, error)
}

type ProductService interface {
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	GetProducts(ctx context.Context, offset int, limit int) ([]entity.Product, error)
}
