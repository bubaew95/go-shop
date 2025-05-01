package postgresql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/bubaew95/go_shop/conf"
	"github.com/bubaew95/go_shop/internal/application/product/entity"
	"github.com/bubaew95/go_shop/pkg/helpers"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func initRepository() (*ProductRepository, error) {
	config := conf.DatabaseConfig{
		Driver: "pgx",
		Dsn:    "host=127.0.0.1 user=shop password=admin dbname=shop sslmode=disable",
	}

	database, err := helpers.NewDB(&config)

	return NewProductRepository(database), err
}

func TestProductRepository_CreateProduct(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		data entity.Product
		want entity.Product
	}{
		{
			name: "success",
			data: entity.Product{
				FirmID:         1,
				UserID:         1,
				Name:           "iPhone 16 Pro Max",
				Anons:          "test annons",
				Text:           "test text",
				Stock:          1,
				Price:          120.5,
				Discount:       10,
				SeoTitle:       "test seo title",
				SeoDescription: "test seo description",
				SeoKeywords:    "test seo keywords",
			},
			want: entity.Product{
				Name:  "iPhone 16 Pro Max",
				Price: 120.5,
			},
		},
	}

	productRepo, err := initRepository()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := productRepo.CreateProduct(context.Background(), tt.data)
			require.NoError(t, err)

			assert.Equal(t, tt.want.Name, product.Name)
			assert.Equal(t, tt.want.Price, product.Price)
		})
	}
}

func BenchmarkProductRepository_GetProducts(b *testing.B) {
	productRepo, _ := initRepository()

	b.Run("N+1 queries", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			productRepo.GetProducts(context.Background(), 0, 10)
		}
	})

	b.Run("Use lateral", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			productRepo.GetProductsSecondVersion(context.Background(), 0, 10)
		}
	})
}
