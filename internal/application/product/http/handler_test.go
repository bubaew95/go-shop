package http

import (
	"encoding/json"
	"fmt"
	"github.com/bubaew95/go_shop/internal/application/product/entity"
	"github.com/bubaew95/go_shop/internal/application/product/infra/postgresql/mock"
	"github.com/bubaew95/go_shop/internal/application/product/service"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	t.Parallel()

	type want struct {
		statusCode  int
		contentType string
		response    string
	}

	tests := []struct {
		name string
		data string
		want want
		err  error
	}{
		{
			name: "success",
			data: `{ "name": "test", "price": 100.4}`,
			want: want{
				statusCode:  http.StatusCreated,
				contentType: "application/json",
				response:    `{"id":0,"firm_id":null,"user_id":null,"name":"test","price":100.4}`,
			},
			err: nil,
		},
		{
			name: "Validation error",
			data: `{ "name": "" }`,
			want: want{
				statusCode:  http.StatusBadRequest,
				contentType: "application/json",
				response:    `{"name":["is required"],"price":["is required"]}`,
			},
			err: nil,
		},
	}

	route := chi.NewRouter()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productRepo := mock.NewMockProductRepository(ctrl)
	productService := service.NewProductService(productRepo)
	productHandler := NewProductController(productService)

	route.Post("/api/products", productHandler.CreateProduct)

	ts := httptest.NewServer(route)
	t.Cleanup(ts.Close)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var product entity.Product
			err := json.Unmarshal([]byte(tt.data), &product)
			require.NoError(t, err)

			if _, ok := product.Validate(); ok {
				productRepo.EXPECT().
					CreateProduct(gomock.Any(), gomock.Any()).
					Return(product, tt.err)
			}

			req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/products", ts.URL), strings.NewReader(tt.data))
			require.NoError(t, err)

			client := http.Client{}
			res, err := client.Do(req)
			require.NoError(t, err)

			_, err = io.ReadAll(res.Body)
			require.NoError(t, err)

			assert.Equal(t, tt.want.statusCode, res.StatusCode)
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}

}
