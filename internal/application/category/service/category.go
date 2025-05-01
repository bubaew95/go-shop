package service

import (
	"context"

	"github.com/bubaew95/go_shop/internal/application/category/domain"
	"github.com/bubaew95/go_shop/internal/application/category/entity"
)

type CategoryService struct {
	repo domain.CategoryService
}

func NewCategoryService(repo domain.CategoryService) *CategoryService {
	return &CategoryService{repo: repo}
}

func (c CategoryService) Create(ctx context.Context, category entity.Category) (entity.Category, error) {
	return c.repo.Create(ctx, category)
}

func (c CategoryService) GetCategories(ctx context.Context, offset int, limit int) ([]entity.CategoryResponse, error) {
	return c.repo.GetCategories(ctx, offset, limit)
}
