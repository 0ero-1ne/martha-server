package routes

import (
	"server/controllers"
	"server/services"

	"github.com/gin-gonic/gin"
)

var authorService = services.AuthorService{}
var authorController = controllers.NewAuthorController(authorService)

func AuthorRoutes(globalRoute *gin.RouterGroup) {
	routes := globalRoute.Group("/authors")
	routes.GET("/", authorController.GetAll)
	routes.GET("/:id", authorController.GetById)
	routes.POST("/", authorController.Create)
	routes.PUT("/:id", authorController.Update)
	routes.DELETE("/:id", authorController.Delete)

	// many2many author:book
	routes.GET("/:id/books", authorController.GetBooks)
	routes.POST("/:id/books/:book_id", authorController.AddBook)
	routes.DELETE("/:id/books/:book_id", authorController.DeleteBook)
}
