package service

import (
	"context"

	ports "github.com/bubaew95/go_shop/internal/application/product/domain"
	"github.com/bubaew95/go_shop/internal/application/product/entity"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(r ports.ProductRepository) *ProductService {
	return &ProductService{
		repo: r,
	}
}

func (p ProductService) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	product, err := p.repo.CreateProduct(ctx, product)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p ProductService) GetProducts(ctx context.Context, offset int, limit int) ([]entity.ProductResponse, error) {
	products, err := p.repo.GetProducts(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	return products, nil
}
