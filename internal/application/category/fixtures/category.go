package fixtures

import (
	"github.com/bubaew95/go_shop/internal/application/category/entity"
	"time"

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

func GenerateCategory(n int) []entity.Category {
	var list []entity.Category
	for i := 0; i < n; i++ {
		list = append(list, FakeCategory())
	}
	return list
}
