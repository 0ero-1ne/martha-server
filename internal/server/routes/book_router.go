package routes

import (
	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func bookRouter(globalRoute *gin.RouterGroup, controller controllers.BookController) {
	routes := globalRoute.Group("/books")
	routes.GET("/", controller.GetAll)
	routes.GET("/:book_id", middlewares.ParseParamsId([]string{"book_id"}), controller.GetById)
	routes.POST("/", controller.Create)
	routes.PUT("/:book_id", middlewares.ParseParamsId([]string{"book_id"}), controller.Update)
	routes.DELETE("/:book_id", middlewares.ParseParamsId([]string{"book_id"}), controller.Delete)

	// many2many book:tag
	routes.GET(
		"/:book_id/tags",
		middlewares.ParseParamsId([]string{"book_id"}),
		controller.GetTags)

	routes.POST(
		"/:book_id/tags/:tag_id",
		middlewares.ParseParamsId([]string{"book_id", "tag_id"}),
		controller.AddTag)

	routes.DELETE(
		"/:book_id/tags/:tag_id",
		middlewares.ParseParamsId([]string{"book_id", "tag_id"}),
		controller.DeleteTag)

	// many2many book:author
	routes.GET(
		"/:book_id/authors",
		middlewares.ParseParamsId([]string{"book_id"}),
		controller.GetAuthors)

	routes.POST(
		"/:book_id/authors/:author_id",
		middlewares.ParseParamsId([]string{"book_id", "author_id"}),
		controller.AddAuthor)

	routes.DELETE(
		"/:book_id/authors/:author_id",
		middlewares.ParseParamsId([]string{"book_id", "author_id"}),
		controller.DeleteAuthor)
}
