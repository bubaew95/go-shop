package main

import (
	"fmt"
	"github.com/bubaew95/go_shop/conf"
	"github.com/bubaew95/go_shop/internal/application/product/http"
	"github.com/bubaew95/go_shop/internal/application/product/infra/postgresql"
	"github.com/bubaew95/go_shop/internal/application/product/service"
	"github.com/bubaew95/go_shop/internal/infra/logger"
	"github.com/bubaew95/go_shop/internal/infra/server"
	"github.com/bubaew95/go_shop/pkg/helpers"
	"github.com/go-chi/chi/v5"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	productRepo := postgresql.NewProductRepository(database)
	productService := service.NewProductService(productRepo)
	productHandler := http.NewProductController(productService)

	route := chi.NewRouter()

	route.Route("/products", func(r chi.Router) {
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
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
