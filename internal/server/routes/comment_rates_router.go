package routes

import (
	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/middlewares"
	"github.com/0ero-1ne/martha-server/internal/utils"
	"github.com/gin-gonic/gin"
)

func commentRateRouter(globalRoute *gin.RouterGroup, controller controllers.CommentRateController, jwtManager utils.JWTManager) {
	routes := globalRoute.Group("comment_rates")

	routes.GET("", controller.GetAll)
	routes.POST("",
		middlewares.IsAuth(jwtManager),
		controller.Create)
	routes.PUT("",
		middlewares.IsAuth(jwtManager),
		controller.Update)
	routes.DELETE("/:comment_id/:comment_user_id",
		middlewares.IsAuth(jwtManager),
		middlewares.ParseParamsId([]string{"comment_id", "comment_user_id"}),
		controller.Delete)
}
