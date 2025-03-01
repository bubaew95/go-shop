package conroller

import (
	"encoding/json"
	"github.com/bubaew95/go_shop/internal/adapter/logger"
	entity "github.com/bubaew95/go_shop/internal/core/entity/repository"
	"github.com/bubaew95/go_shop/internal/core/entity/response"
	ports "github.com/bubaew95/go_shop/internal/core/ports/repository/postgresql"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type Product struct {
	router  *chi.Mux
	service ports.ProductService
}

func NewProductController(r *chi.Mux, s ports.ProductService) *Product {
	return &Product{
		router:  r,
		service: s,
	}
}

func (p Product) InitRoute() {
	p.router.Route("/products", func(r chi.Router) {
		r.Post("/", p.CreateProduct)
		r.Get("/", p.GetProducts)
	})
}

func (p Product) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		logger.Log.Debug("Error decoding product", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := p.service.CreateProduct(r.Context(), product)
	if err != nil {
		logger.Log.Debug("Error creating product", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	WriteJson(w, res, http.StatusCreated)
}

func (p Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	offset, limit := ParsePaginate(r)

	products, err := p.service.GetProducts(r.Context(), offset, limit)
	if err != nil {
		logger.Log.Error("Get products error", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
	}

	responseDTO := response.ResponseWithPagination{
		Items:  products,
		Offset: offset,
		Limit:  limit,
	}

	WriteJson(w, responseDTO, http.StatusOK)
}
