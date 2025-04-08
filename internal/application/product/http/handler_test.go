package http

import (
	"github.com/bubaew95/go_shop/internal/application/product/entity"
	"net/http"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	t.Parallel()

	type want struct {
		statusCode  int
		contentType string
		response    entity.Product
	}

	tests := []struct {
		name string
		data string
		want want
	}{
		{
			name: "success",
			data: `{}`,
			want: want{
				statusCode:  http.StatusCreated,
				contentType: "application/json",
				response:    entity.Product{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}

}
