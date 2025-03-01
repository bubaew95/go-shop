package postgresql

import (
	"context"
	"github.com/bubaew95/go_shop/internal/adapter/logger"
	entity "github.com/bubaew95/go_shop/internal/core/entity/repository"
	"github.com/bubaew95/go_shop/internal/infra/repository"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

const tableName = "product"

type ProductRepository struct {
	db *repository.DataBase
}

func NewProductRepository(db *repository.DataBase) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (p ProductRepository) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	sqlQuery := `INSERT INTO ` + tableName + ` (name,  text, price) ` +
		`VALUES ($1, $2, $3) RETURNING id`
	logger.Log.Debug("Sql create product",
		zap.String("sql", sqlQuery),
		zap.Any("product", product),
	)

	var lastId int
	err := p.db.QueryRowContext(ctx, sqlQuery,
		product.Name,
		product.TEXT,
		product.Price,
	).Scan(&lastId)

	if err != nil {
		return entity.Product{}, err
	}

	product = entity.Product{
		ID:    lastId,
		Name:  product.Name,
		TEXT:  product.TEXT,
		Price: product.Price,
	}

	return product, nil
}

func (p ProductRepository) GetProducts(ctx context.Context, offset int, limit int) ([]entity.Product, error) {
	sqlQuery := `SELECT id, name, price, anons, text, sale, active  FROM ` + tableName +
		` WHERE active = 1 LIMIT $1 OFFSET $2`

	logger.Log.Debug("Sql get products",
		zap.String("sql", sqlQuery),
		zap.Int("offset", offset),
		zap.Int("limit", limit),
	)

	rows, err := p.db.QueryContext(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Anons, &product.TEXT, &product.Sale, &product.Active)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
