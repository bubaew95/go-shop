package logger

import (
	"go.uber.org/zap"

	"github.com/bubaew95/go_shop/conf"
)

var Log *zap.Logger = zap.NewNop()

func Load(c *conf.ServerConfig) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	if c.Debug {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return err
		}
	}

	Log = logger

	return nil
}
