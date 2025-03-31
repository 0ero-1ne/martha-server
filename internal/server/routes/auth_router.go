package routes

import (
	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/gin-gonic/gin"
)

func authRouter(globalRoute *gin.RouterGroup, controller controllers.AuthController) {
	router := globalRoute.Group("/auth")
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	router.POST("/refresh", controller.Refresh)
}
