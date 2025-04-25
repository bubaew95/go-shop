package helpers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/bubaew95/go_shop/internal/infra/logger"
)

func WriteJson(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Log.Error("Json encode error", zap.Error(err))
	}
}

func ParsePaginate(r *http.Request) (offset int, limit int) {
	var (
		defaultOffset = 0
		defaultLimit  = 20
		maxLimit      = 100
	)

	offsetQuery := r.URL.Query().Get("offset")
	limitQuery := r.URL.Query().Get("limit")

	offset, err := safeAtoi(offsetQuery, defaultOffset)
	if err != nil || offset < 0 {
		return defaultOffset, defaultLimit
	}

	limit, err = safeAtoi(limitQuery, defaultLimit)
	if err != nil || limit <= 0 || limit > maxLimit {
		return defaultOffset, defaultLimit
	}

	return offset, limit
}

func safeAtoi(v string, defaultValue int) (int, error) {
	if v == "" {
		return defaultValue, nil
	}

	result, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return result, nil
}
