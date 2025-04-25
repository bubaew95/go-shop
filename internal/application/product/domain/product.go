package domain

import (
	"context"

	"github.com/bubaew95/go_shop/internal/application/product/entity"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	GetProducts(ctx context.Context, offset int, limit int) ([]entity.ProductResponse, error)
}

type ProductService interface {
	CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	GetProducts(ctx context.Context, offset int, limit int) ([]entity.ProductResponse, error)
}
