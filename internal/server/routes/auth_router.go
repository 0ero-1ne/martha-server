package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/0ero-1ne/martha-server/internal/controllers"
)

func authRouter(globalRoute *gin.RouterGroup, controller controllers.AuthController) {
	router := globalRoute.Group("/auth")
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	router.POST("/refresh", controller.Refresh)
}
