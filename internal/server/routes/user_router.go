package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/middlewares"
	"github.com/0ero-1ne/martha-server/internal/utils"
)

func userRouter(
	globalRoute *gin.RouterGroup,
	controller controllers.UserController,
	jwtManager utils.JWTManager,
) {
	router := globalRoute.Group("/users")
	router.GET("", controller.GetAll)
	router.GET("/make_moderator/:user_id", middlewares.ParseParamsId([]string{"user_id"}), controller.MakeModer)
	router.GET("/make_user/:user_id", middlewares.ParseParamsId([]string{"user_id"}), controller.MakeUser)
	router.GET("/count", controller.GetCount)
	router.GET("/single", middlewares.IsAuth(jwtManager), controller.GetById)
	router.POST("", middlewares.IsAuth(jwtManager), controller.Update)
	router.DELETE("/:user_id", middlewares.ParseParamsId([]string{"user_id"}), controller.Delete)
}
