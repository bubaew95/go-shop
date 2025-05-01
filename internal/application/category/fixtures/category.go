package fixtures

import (
	"context"
	"fmt"
	"time"

	"github.com/bubaew95/go_shop/internal/application/category/entity"
	"github.com/bubaew95/go_shop/internal/application/category/infra/postgresql"
	"github.com/bubaew95/go_shop/pkg/helpers"

	"github.com/brianvoe/gofakeit/v6"
)

func init() {
	gofakeit.Seed(0)
}

func FakeCategory() entity.Category {
	return entity.Category{
		Name:      gofakeit.Name(),
		Alias:     gofakeit.EmojiAlias(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func GenerateCategory(n int, db *helpers.DataBase) {
	repo := postgresql.NewCategoryRepository(db)

	for i := 0; i < n; i++ {
		category := FakeCategory()

		_, err := repo.Create(context.Background(), category)
		if err != nil {
			fmt.Println(err)
		}
	}
}
