package routes

import (
	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/middlewares"
	"github.com/0ero-1ne/martha-server/internal/utils"
	"github.com/gin-gonic/gin"
)

func bookRateRouter(globalRoute *gin.RouterGroup, controller controllers.BookRateController, jwtManager utils.JWTManager) {
	routes := globalRoute.Group("book_rates")

	routes.GET("", controller.GetAll)
	routes.POST("",
		middlewares.IsAuth(jwtManager),
		controller.Create)
	routes.PUT("",
		middlewares.IsAuth(jwtManager),
		controller.Update)
	routes.DELETE("/:book_id/:book_user_id",
		middlewares.IsAuth(jwtManager),
		middlewares.ParseParamsId([]string{"book_id", "book_user_id"}),
		controller.Delete)
}
