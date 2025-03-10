package routes

import (
	"server/controllers"
	"server/services"

	"github.com/gin-gonic/gin"
)

var service services.TagService = services.NewTagService()
var controller controllers.TagController = controllers.NewTagController(service)

func TagRoutes(globalRoute *gin.RouterGroup) {
	routes := globalRoute.Group("/tags")
	routes.GET("/", controller.GetAll)
	routes.GET("/:id", controller.GetById)
	routes.POST("/", controller.Create)
	routes.PUT("/:id", controller.Update)
	routes.DELETE("/:id", controller.Delete)
}
