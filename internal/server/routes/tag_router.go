package routes

import (
	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func tagRouter(globalRoute *gin.RouterGroup, controller controllers.TagController) {
	routes := globalRoute.Group("/tags")
	routes.GET("/", controller.GetAll)
	routes.GET("/:tag_id", middlewares.ParseParamsId([]string{"tag_id"}), controller.GetById)
	routes.POST("/", controller.Create)
	routes.PUT("/:tag_id", middlewares.ParseParamsId([]string{"tag_id"}), controller.Update)
	routes.DELETE("/:tag_id", middlewares.ParseParamsId([]string{"tag_id"}), controller.Delete)

	// many2many tag:book
	routes.GET(
		"/:tag_id/books",
		middlewares.ParseParamsId([]string{"tag_id"}),
		controller.GetBooks)

	routes.POST(
		"/:tag_id/books/:book_id",
		middlewares.ParseParamsId([]string{"tag_id", "book_id"}),
		controller.AddBook)

	routes.DELETE(
		"/:tag_id/books/:book_id",
		middlewares.ParseParamsId([]string{"tag_id", "book_id"}),
		controller.DeleteBook)
}
