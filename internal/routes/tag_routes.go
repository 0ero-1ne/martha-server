package routes

import (
	"github.com/0ero-1ne/martha/internal/controllers"
	"github.com/0ero-1ne/martha/internal/middlewares"
	"github.com/0ero-1ne/martha/internal/services"

	"github.com/gin-gonic/gin"
)

var tagService = services.NewTagService()
var tagController = controllers.NewTagController(tagService)

func TagRoutes(globalRoute *gin.RouterGroup) {
	routes := globalRoute.Group("/tags")
	routes.GET("/", tagController.GetAll)
	routes.GET("/:tag_id", middlewares.ParseParamsId([]string{"tag_id"}), tagController.GetById)
	routes.POST("/", tagController.Create)
	routes.PUT("/:tag_id", middlewares.ParseParamsId([]string{"tag_id"}), tagController.Update)
	routes.DELETE("/:tag_id", middlewares.ParseParamsId([]string{"tag_id"}), tagController.Delete)

	// many2many tag:book
	routes.GET(
		"/:tag_id/books",
		middlewares.ParseParamsId([]string{"tag_id"}),
		tagController.GetBooks)

	routes.POST(
		"/:tag_id/books/:book_id",
		middlewares.ParseParamsId([]string{"tag_id", "book_id"}),
		tagController.AddBook)

	routes.DELETE(
		"/:tag_id/books/:book_id",
		middlewares.ParseParamsId([]string{"tag_id", "book_id"}),
		tagController.DeleteBook)
}
