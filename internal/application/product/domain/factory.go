package domain

import "github.com/bubaew95/go_shop/internal/application/product/entity"

type Factory interface {
	Create(product entity.Product) (*entity.Product, error)
}

type DefaultFactory struct{}

func NewDefaultFactory() *DefaultFactory {
	return &DefaultFactory{}
}

func (f *DefaultFactory) Create(product entity.Product) (*entity.Product, error) {

	return &product, nil
}
