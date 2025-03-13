package routes

import (
	"server/controllers"
	"server/services"

	"github.com/gin-gonic/gin"
)

var tagService = services.NewTagService()
var tagController = controllers.NewTagController(tagService)

func TagRoutes(globalRoute *gin.RouterGroup) {
	routes := globalRoute.Group("/tags")
	routes.GET("/", tagController.GetAll)
	routes.GET("/:id", tagController.GetById)
	routes.POST("/", tagController.Create)
	routes.PUT("/:id", tagController.Update)
	routes.DELETE("/:id", tagController.Delete)

	// many2many tag:book
	routes.GET("/:id/books", tagController.GetBooks)
	routes.POST("/:id/books/:book_id", tagController.AddBook)
	routes.DELETE("/:id/books/:book_id", tagController.DeleteBook)
}
