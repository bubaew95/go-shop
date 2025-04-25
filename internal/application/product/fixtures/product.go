package fixtures

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/bubaew95/go_shop/internal/application/product/entity"
	"github.com/bubaew95/go_shop/internal/application/product/infra/postgresql"
	"github.com/bubaew95/go_shop/pkg/helpers"
)

func init() {
	gofakeit.Seed(0)
}

func NewFakeProduct() entity.Product {
	product := gofakeit.Product()

	return entity.Product{
		Name:           product.Name,
		FirmID:         gofakeit.Number(1, 10),
		UserID:         gofakeit.Number(1, 10),
		Anons:          product.Description,
		Price:          product.Price,
		Text:           gofakeit.Sentence(1024),
		Stock:          gofakeit.Number(0, 10),
		Discount:       gofakeit.Number(0, 10),
		SeoTitle:       gofakeit.Sentence(3),
		SeoKeywords:    gofakeit.Sentence(3),
		SeoDescription: gofakeit.Sentence(3),
		CreatedAt:      gofakeit.Date(),
		UpdatedAt:      gofakeit.Date(),
	}
}

func GenerateProductFixtures(n int, db *helpers.DataBase) {
	repo := postgresql.NewProductRepository(db)

	for i := 0; i < n; i++ {
		_, err := repo.CreateProduct(context.Background(), NewFakeProduct())
		if err != nil {
			fmt.Println(err)
		}
	}
}
