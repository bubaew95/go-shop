package postgresql

import (
	"context"
	"encoding/json"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	"github.com/bubaew95/go_shop/internal/application/product/entity"
	"github.com/bubaew95/go_shop/internal/infra/logger"
	"github.com/bubaew95/go_shop/pkg/helpers"
)

const tableName = "product"

type ProductRepository struct {
	db *helpers.DataBase
}

func NewProductRepository(db *helpers.DataBase) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// CreateProduct - добавление товара в базу
func (p ProductRepository) CreateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	sqlQuery := `INSERT INTO product (firm_id, user_id, name, anons, text, stock, price, discount, seo_title, seo_description, seo_keywords, created_at, updated_at) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`

	logger.Log.Debug("Sql create product",
		zap.String("sql", sqlQuery),
		zap.Any("product", product),
	)

	var lastId int
	err := p.db.QueryRowContext(ctx, sqlQuery,
		product.FirmID,
		product.UserID,
		product.Name,
		product.Anons,
		product.Text,
		product.Stock,
		product.Price,
		product.Discount,
		product.SeoTitle,
		product.SeoDescription,
		product.SeoKeywords,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(&lastId)

	if err != nil {
		return entity.Product{}, err
	}

	product = entity.Product{
		ID:    lastId,
		Name:  product.Name,
		Text:  product.Text,
		Price: product.Price,
	}

	return product, nil
}

func (p ProductRepository) GetProducts(ctx context.Context, offset int, limit int) ([]entity.ProductResponse, error) {
	sqlQuery := `
		SELECT p.id, p.name,p.price, p.firm_id
		FROM product p
        ORDER BY p.id DESC
        LIMIT  $1 OFFSET $2
	`

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

	var products []entity.ProductResponse
	for rows.Next() {
		var product entity.ProductResponse
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.FirmID)
		if err != nil {
			return nil, err
		}

		categories, err := p.GetCategoriesByProductID(ctx, product.ID)
		if err == nil {
			return nil, err
		}

		product.Category = categories

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (p ProductRepository) GetCategoriesByProductID(ctx context.Context, producID int) ([]entity.Category, error) {
	sqlQuery := `
		SELECT c.id, c.name FROM category c
		LEFT JOIN categories_to_product ctp ON c.id = ctp.category_id
		WHERE ctp.product_id = $1
	`

	rows, err := p.db.QueryContext(ctx, sqlQuery, producID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (p ProductRepository) GetProductsSecondVersion(ctx context.Context, offset int, limit int) ([]entity.ProductResponse, error) {
	sqlQuery := `
		SELECT 
            p.id, 
            p.name,
            p.price,
            p.firm_id,
            COALESCE(cat.categories, '[]') AS category
        FROM product p
        LEFT JOIN LATERAL (
            SELECT 
              json_agg(json_build_object('id', c.id, 'name', c.name)) AS categories
            FROM categories_to_product ctp
            JOIN category c ON c.id = ctp.category_id
            WHERE ctp.product_id = p.id
        ) AS cat ON TRUE
        ORDER BY p.id DESC
        LIMIT  $1
        OFFSET $2
	`

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

	var products []entity.ProductResponse
	for rows.Next() {
		var product entity.ProductResponse
		var catJSON []byte

		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.FirmID, &catJSON)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(catJSON, &product.Category); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
