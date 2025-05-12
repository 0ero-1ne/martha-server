package routes

import (
	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/middlewares"
	"github.com/0ero-1ne/martha-server/internal/utils"
	"github.com/gin-gonic/gin"
)

func commentRouter(globalRoute *gin.RouterGroup, controller controllers.CommentController, jwtManager utils.JWTManager) {
	routes := globalRoute.Group("/comments")

	routes.GET("", controller.GetAll)
	routes.GET("/:comment_id",
		middlewares.ParseParamsId([]string{"comment_id"}),
		controller.GetById)
	routes.POST("",
		middlewares.IsAuth(jwtManager),
		controller.Create)
	routes.PUT("/:comment_id",
		middlewares.ParseParamsId([]string{"comment_id"}),
		controller.Update)
	routes.DELETE("/:comment_id",
		middlewares.ParseParamsId([]string{"author_id"}),
		controller.Delete)

	routes.GET("/book/:book_id",
		middlewares.ParseParamsId([]string{"book_id"}),
		controller.GetAllByBookId)
}
