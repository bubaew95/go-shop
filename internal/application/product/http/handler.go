package http

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	ports "github.com/bubaew95/go_shop/internal/application/product/domain"
	"github.com/bubaew95/go_shop/internal/application/product/entity"
	"github.com/bubaew95/go_shop/internal/infra/logger"
	"github.com/bubaew95/go_shop/pkg/helpers"
	"github.com/bubaew95/go_shop/pkg/model/response"
)

type ProductHandler struct {
	service ports.ProductService
}

func NewProductController(s ports.ProductService) *ProductHandler {
	return &ProductHandler{
		service: s,
	}
}

func (p ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		logger.Log.Debug("Error decoding product", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errs, ok := product.Validate()
	if !ok {
		helpers.WriteJson(w, errs, http.StatusBadRequest)
		return
	}

	res, err := p.service.CreateProduct(r.Context(), product)
	if err != nil {
		logger.Log.Debug("Error creating product", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.WriteJson(w, res, http.StatusCreated)
}

func (p ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	offset, limit := helpers.ParsePaginate(r)

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

	helpers.WriteJson(w, responseDTO, http.StatusOK)
}
