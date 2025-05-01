package http

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bubaew95/go_shop/internal/application/category/entity"
	"github.com/bubaew95/go_shop/internal/application/category/infra/postgresql/mock"
	categoryS "github.com/bubaew95/go_shop/internal/application/category/service"
	"github.com/bubaew95/go_shop/pkg/model/response"
)

func initMockCategoryServer(t *testing.T) (*mock.MockCategoryRepository, *httptest.Server) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mock.NewMockCategoryRepository(ctrl)
	service := categoryS.NewCategoryService(repo)
	handler := NewCategoryController(service)

	route := chi.NewRouter()
	route.Route("/api/category", func(r chi.Router) {
		r.Get("/", handler.GetCategories)
	})

	ts := httptest.NewServer(route)
	t.Cleanup(ts.Close)

	return repo, ts
}

func TestCategoryHandler_GetCategories(t *testing.T) {
	t.Parallel()

	type want struct {
		contentType string
		statusCode  int
		data        response.ResponseWithPagination
	}

	type db struct {
		response []entity.CategoryResponse
		err      error
	}

	categories := []entity.CategoryResponse{
		{
			ID:    1,
			Name:  "Тестовая категория",
			Alias: "test",
		},
	}

	tests := []struct {
		name string
		want want
		db   db
	}{
		{
			name: "Get all categories",
			db: db{
				response: categories,
				err:      nil,
			},
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusOK,
				data: response.ResponseWithPagination{
					Items:  categories,
					Offset: 0,
					Limit:  1,
				},
			},
		},
		{
			name: "Categories empty",
			db: db{
				response: []entity.CategoryResponse{},
				err:      nil,
			},
			want: want{
				contentType: "application/json",
				statusCode:  http.StatusOK,
				data: response.ResponseWithPagination{
					Items:  []entity.CategoryResponse{},
					Offset: 0,
					Limit:  1,
				},
			},
		},
	}

	repo, ts := initMockCategoryServer(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.EXPECT().
				GetCategories(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.db.response, tt.db.err)

			req, err := http.NewRequest(http.MethodGet, ts.URL+"/api/category", nil)
			require.NoError(t, err)

			q := req.URL.Query()
			q.Add("offset", strconv.Itoa(tt.want.data.Offset))
			q.Add("limit", strconv.Itoa(tt.want.data.Limit))
			req.URL.RawQuery = q.Encode()

			client := http.Client{}
			res, err := client.Do(req)
			require.NoError(t, err)

			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			defer res.Body.Close()

			responseBody, err := json.Marshal(tt.want.data)
			require.NoError(t, err)

			assert.JSONEq(t, string(responseBody), string(resBody))
			assert.Equal(t, tt.want.statusCode, res.StatusCode)
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
			assert.Equal(t, strconv.Itoa(tt.want.data.Limit), req.URL.Query().Get("limit"))
		})
	}
}
