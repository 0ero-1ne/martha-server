package routes

import (
	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/middlewares"
	"github.com/0ero-1ne/martha-server/internal/utils"
	"github.com/gin-gonic/gin"
)

func userRouter(globalRoute *gin.RouterGroup, controller controllers.UserController, jwtManager utils.JWTManager) {
	router := globalRoute.Group("/users")
	router.GET("/single", middlewares.IsAuth(jwtManager), controller.GetById)
}
