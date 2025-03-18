package app

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/0ero-1ne/martha-server/internal/config"
	"github.com/0ero-1ne/martha-server/internal/db"
	"github.com/0ero-1ne/martha-server/internal/db/postgres"
	"github.com/0ero-1ne/martha-server/internal/server"
	"github.com/0ero-1ne/martha-server/internal/server/routes"
	"github.com/0ero-1ne/martha-server/internal/utils"
)

func Run(configPath string) {
	cfg := config.Init(configPath)
	postgresDB := postgres.NewPostgresDatabase(cfg.PostgresConfig)

	database, err := db.InitDatabase(postgresDB)
	if err != nil {
		panic("Can not connect to database: " + err.Error())
	}

	jwtManager := utils.NewJWTManager(cfg.JWTConfig)
	httpServer := server.NewHttpServer(cfg.ServerConfig, routes.NewRouter(database, jwtManager))

	go func() {
		log.Printf("Sever is listening on %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("Can not init server: " + err.Error())
		}
	}()

	done := make(chan bool, 1)

	go gracefulShutdown(httpServer, done)

	<-done
	log.Println("Graceful shutdown complete.")
}

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
	done <- true
}
