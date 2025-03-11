package routes

import (
	"server/controllers"
	"server/services"

	"github.com/gin-gonic/gin"
)

var bookService = services.NewBookService()
var bookController = controllers.NewBookController(bookService)

func BookRoutes(globalRoute *gin.RouterGroup) {
	routes := globalRoute.Group("/books")
	routes.GET("/", bookController.GetAll)
	routes.GET("/:id", bookController.GetById)
	routes.POST("/", bookController.Create)
	routes.PUT("/:id", bookController.Update)
	routes.DELETE("/:id", bookController.Delete)

	// many2many book:tag
	routes.POST("/:id/tags/:tag_id", bookController.AddTag)
	routes.DELETE("/:id/tags/:tag_id", bookController.DeleteTag)
}
