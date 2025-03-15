package routes

import (
	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func authorRouter(globalRoute *gin.RouterGroup, controller controllers.AuthorController) {
	routes := globalRoute.Group("/authors")
	routes.GET("/", controller.GetAll)
	routes.GET("/:author_id", middlewares.ParseParamsId([]string{"author_id"}), controller.GetById)
	routes.POST("/", controller.Create)
	routes.PUT("/:author_id", middlewares.ParseParamsId([]string{"author_id"}), controller.Update)
	routes.DELETE("/:author_id", middlewares.ParseParamsId([]string{"author_id"}), controller.Delete)

	// many2many author:book
	routes.GET(
		"/:author_id/books",
		middlewares.ParseParamsId([]string{"author_id"}),
		controller.GetBooks)

	routes.POST(
		"/:author_id/books/:book_id",
		middlewares.ParseParamsId([]string{"author_id", "book_id"}),
		controller.AddBook)

	routes.DELETE(
		"/:author_id/books/:book_id",
		middlewares.ParseParamsId([]string{"author_id", "book_id"}),
		controller.DeleteBook)
}
