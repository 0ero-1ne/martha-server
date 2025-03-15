package app

import (
	"github.com/0ero-1ne/martha/internal/config"
	"github.com/0ero-1ne/martha/internal/routes"

	"github.com/0ero-1ne/martha/internal/db"

	"github.com/gin-gonic/gin"
)

func Run(configPath string) {
	cfg := config.Init(configPath)
	db.InitDatabase(cfg.PostgresConfig)

	server := gin.Default()

	globalRoute := server.Group("/api/v1")
	routes.TagRoutes(globalRoute)
	routes.BookRoutes(globalRoute)
	routes.AuthorRoutes(globalRoute)
	routes.ChapterRoutes(globalRoute)

	err := server.Run(cfg.ServerConfig.GetAddress())

	if err != nil {
		panic("Can not start server: " + err.Error())
	}
}
