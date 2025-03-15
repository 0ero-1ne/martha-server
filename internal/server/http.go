package server

import (
	"net/http"
	"time"

	"github.com/0ero-1ne/martha-server/internal/config"
)

func NewHttpServer(config config.ServerConfig, router http.Handler) *http.Server {
	httpServer := http.Server{
		Addr:           config.GetAddress(),
		Handler:        router,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	return &httpServer
}
