package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/bubaew95/go_shop/conf"
)

const defaultHost = "0.0.0.0"

type HttpServer interface {
	Start()
	Stop()
}

type httpServer struct {
	Port   uint
	server *http.Server
}

func NewHttpServer(router *chi.Mux, c *conf.ServerConfig) HttpServer {
	return &httpServer{
		Port: c.Port,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", defaultHost, c.Port),
			Handler: router,
		},
	}
}

func (httpServer httpServer) Start() {
	go func() {
		if err := httpServer.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf(
				"failed to start HttpServer on port %d, err=%s",
				httpServer.Port, err.Error(),
			)
		}
	}()

	fmt.Println("http server started")
}

func (httpServer httpServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3)*time.Second)
	defer cancel()

	if err := httpServer.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown, err=%s", err.Error())
	}
}
