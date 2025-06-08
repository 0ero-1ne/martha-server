package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/middlewares"
	"github.com/0ero-1ne/martha-server/internal/utils"
)

func authRouter(globalRoute *gin.RouterGroup, controller controllers.AuthController, jwtManager utils.JWTManager) {
	router := globalRoute.Group("/auth")
	router.POST("/signup", controller.Signup)
	router.POST("/login", controller.Login)
	router.POST("/refresh", controller.Refresh)
	router.POST("/change_password", middlewares.IsAuth(jwtManager), controller.ChangePassword)
}
