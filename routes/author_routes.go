package routes

import (
	"server/controllers"
	"server/middlewares"
	"server/services"

	"github.com/gin-gonic/gin"
)

var authorService = services.AuthorService{}
var authorController = controllers.NewAuthorController(authorService)

func AuthorRoutes(globalRoute *gin.RouterGroup) {
	routes := globalRoute.Group("/authors")
	routes.GET("/", authorController.GetAll)
	routes.GET("/:author_id", middlewares.ParseParamsId([]string{"author_id"}), authorController.GetById)
	routes.POST("/", authorController.Create)
	routes.PUT("/:author_id", middlewares.ParseParamsId([]string{"author_id"}), authorController.Update)
	routes.DELETE("/:author_id", middlewares.ParseParamsId([]string{"author_id"}), authorController.Delete)

	// many2many author:book
	routes.GET(
		"/:author_id/books",
		middlewares.ParseParamsId([]string{"author_id"}),
		authorController.GetBooks)

	routes.POST(
		"/:author_id/books/:book_id",
		middlewares.ParseParamsId([]string{"author_id", "book_id"}),
		authorController.AddBook)

	routes.DELETE(
		"/:author_id/books/:book_id",
		middlewares.ParseParamsId([]string{"author_id", "book_id"}),
		authorController.DeleteBook)
}
