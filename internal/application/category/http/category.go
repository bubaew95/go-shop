package http

import (
	"github.com/bubaew95/go_shop/internal/infra/logger"
	"github.com/bubaew95/go_shop/pkg/helpers"
	"github.com/bubaew95/go_shop/pkg/model/response"
	"go.uber.org/zap"
	"net/http"

	"github.com/bubaew95/go_shop/internal/application/category/domain"
)

type CategoryHandler struct {
	service domain.CategoryService
}

func NewCategoryController(s domain.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: s,
	}
}

func (c CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	offset, limit := helpers.ParsePaginate(r)

	categories, err := c.service.GetCategories(r.Context(), offset, limit)
	if err != nil {
		logger.Log.Error("Get categories error", zap.Error(err))
		w.WriteHeader(http.StatusNoContent)
	}

	responseDTO := response.ResponseWithPagination{
		Items:  categories,
		Offset: offset,
		Limit:  limit,
	}

	helpers.WriteJson(w, responseDTO, http.StatusOK)
}
