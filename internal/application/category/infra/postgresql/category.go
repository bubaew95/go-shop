package postgresql

import (
	"context"

	"github.com/bubaew95/go_shop/internal/application/category/entity"
	"github.com/bubaew95/go_shop/pkg/helpers"
)

type CategoryRepository struct {
	db *helpers.DataBase
}

func NewCategoryRepository(db *helpers.DataBase) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) Create(ctx context.Context, category entity.Category) (entity.Category, error) {
	sqlQuery := `INSERT INTO category (name, alias) VALUES ($1, $2) RETURNING id`

	var lastId int
	err := repo.db.QueryRowContext(ctx, sqlQuery, category.Name, category.Alias).Scan(&lastId)
	if err != nil {
		return entity.Category{}, err
	}

	return entity.Category{
		ID:    lastId,
		Name:  category.Name,
		Alias: category.Alias,
	}, nil
}

func (repo *CategoryRepository) GetCategories(ctx context.Context, offset int, limit int) ([]entity.CategoryResponse, error) {
	sqlQuery := `SELECT id, name, alias FROM category ORDER BY name LIMIT $1 OFFSET $2`

	rows, err := repo.db.QueryContext(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]entity.CategoryResponse, 0)
	for rows.Next() {
		var category entity.CategoryResponse
		if err := rows.Scan(&category.ID, &category.Name, &category.Alias); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
