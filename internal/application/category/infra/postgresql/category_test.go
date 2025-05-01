package postgresql

import (
	"context"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bubaew95/go_shop/conf"
	"github.com/bubaew95/go_shop/pkg/helpers"
)

func TestCategoryRepository_GetCategories(t *testing.T) {
	t.Parallel()

	type request struct {
		offset int
		limit  int
	}

	tests := []struct {
		name string
		data request
	}{
		{
			name: "Get one category",
			data: request{
				offset: 0,
				limit:  1,
			},
		},
		{
			name: "Get two categories",
			data: request{
				offset: 0,
				limit:  2,
			},
		},
	}

	config := conf.DatabaseConfig{
		Driver: "pgx",
		Dsn:    "host=127.0.0.1 user=shop password=admin dbname=shop sslmode=disable",
	}

	database, err := helpers.NewDB(&config)
	require.NoError(t, err)

	repo := NewCategoryRepository(database)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			categories, err := repo.GetCategories(context.Background(), tt.data.offset, tt.data.limit)
			require.NoError(t, err)

			assert.Len(t, categories, tt.data.limit)
			assert.NotEmpty(t, categories)
		})
	}
}
