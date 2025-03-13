package main

import (
	"server/config"
	"server/db"
	"server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Init()
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
