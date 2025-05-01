package main

import (
	"fmt"
	"log"

	"github.com/bubaew95/go_shop/conf"
	categoryFixtures "github.com/bubaew95/go_shop/internal/application/category/fixtures"
	productFixtures "github.com/bubaew95/go_shop/internal/application/product/fixtures"
	"github.com/bubaew95/go_shop/internal/infra/logger"
	"github.com/bubaew95/go_shop/pkg/helpers"
)

func init() {
	if err := conf.LoadEnvOptional(""); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	config := conf.NewServerConfig()
	if err := logger.Load(config); err != nil {
		log.Fatalf("load config failed: %v", err)
	}

	database, err := helpers.NewDB(config.Database)
	if err != nil {
		log.Fatalf("init db failed: %v", err)
	}

	categoryFixtures.GenerateCategory(10, database)
	productFixtures.GenerateProductFixtures(10, database)
}
