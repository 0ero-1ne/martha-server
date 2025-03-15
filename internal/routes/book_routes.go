package routes

import (
	"github.com/0ero-1ne/martha/internal/controllers"
	"github.com/0ero-1ne/martha/internal/middlewares"
	"github.com/0ero-1ne/martha/internal/services"

	"github.com/gin-gonic/gin"
)

var bookService = services.NewBookService()
var bookController = controllers.NewBookController(bookService)

func BookRoutes(globalRoute *gin.RouterGroup) {
	routes := globalRoute.Group("/books")
	routes.GET("/", bookController.GetAll)
	routes.GET("/:book_id", middlewares.ParseParamsId([]string{"book_id"}), bookController.GetById)
	routes.POST("/", bookController.Create)
	routes.PUT("/:book_id", middlewares.ParseParamsId([]string{"book_id"}), bookController.Update)
	routes.DELETE("/:book_id", middlewares.ParseParamsId([]string{"book_id"}), bookController.Delete)

	// many2many book:tag
	routes.GET(
		"/:book_id/tags",
		middlewares.ParseParamsId([]string{"book_id"}),
		bookController.GetTags)

	routes.POST(
		"/:book_id/tags/:tag_id",
		middlewares.ParseParamsId([]string{"book_id", "tag_id"}),
		bookController.AddTag)

	routes.DELETE(
		"/:book_id/tags/:tag_id",
		middlewares.ParseParamsId([]string{"book_id", "tag_id"}),
		bookController.DeleteTag)

	// many2many book:author
	routes.GET(
		"/:book_id/authors",
		middlewares.ParseParamsId([]string{"book_id"}),
		bookController.GetAuthors)

	routes.POST(
		"/:book_id/authors/:author_id",
		middlewares.ParseParamsId([]string{"book_id", "author_id"}),
		bookController.AddAuthor)

	routes.DELETE(
		"/:book_id/authors/:author_id",
		middlewares.ParseParamsId([]string{"book_id", "author_id"}),
		bookController.DeleteAuthor)
}
