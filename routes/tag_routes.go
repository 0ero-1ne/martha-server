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
}
