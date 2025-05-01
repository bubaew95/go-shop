package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	categoryH "github.com/bubaew95/go_shop/internal/application/category/http"
	categoryR "github.com/bubaew95/go_shop/internal/application/category/infra/postgresql"
	categoryS "github.com/bubaew95/go_shop/internal/application/category/service"

	"github.com/go-chi/chi/v5"

	"github.com/bubaew95/go_shop/conf"
	productH "github.com/bubaew95/go_shop/internal/application/product/http"
	productR "github.com/bubaew95/go_shop/internal/application/product/infra/postgresql"
	"github.com/bubaew95/go_shop/internal/application/product/service"
	"github.com/bubaew95/go_shop/internal/infra/logger"
	"github.com/bubaew95/go_shop/internal/infra/server"
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

	productRepo := productR.NewProductRepository(database)
	productService := service.NewProductService(productRepo)
	productHandler := productH.NewProductController(productService)

	categoryRepo := categoryR.NewCategoryRepository(database)
	categoryService := categoryS.NewCategoryService(categoryRepo)
	categoryHandler := categoryH.NewCategoryController(categoryService)

	route := chi.NewRouter()
	route.Route("/products", func(r chi.Router) {
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
	})
	route.Route("/category", func(r chi.Router) {
		r.Get("/", categoryHandler.GetCategories)
	})

	apiRoute := chi.NewRouter()
	apiRoute.Mount("/api", route)

	start(apiRoute, config)
}

func start(apiRoute *chi.Mux, config *conf.ServerConfig) {
	httpServer := server.NewHttpServer(apiRoute, config)
	httpServer.Start()
	defer httpServer.Stop()

	logger.Log.Info("Listening for signal")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-ch
	logger.Log.Info("Graceful shutdown...")
}
