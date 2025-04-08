package routes

import (
	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func chapterRouter(globalRoute *gin.RouterGroup, controller controllers.ChapterController) {
	routes := globalRoute.Group("/chapters")
	routes.GET("/", controller.GetAll)
	routes.GET("/:chapter_id", middlewares.ParseParamsId([]string{"chapter_id"}), controller.GetById)
	routes.GET("/book/:book_id", middlewares.ParseParamsId([]string{"book_id"}), controller.GetChaptersByBookId)
	routes.POST("/", controller.Create)
	routes.PUT("/:chapter_id", middlewares.ParseParamsId([]string{"chapter_id"}), controller.Update)
	routes.DELETE("/:chapter_id", middlewares.ParseParamsId([]string{"chapter_id"}), controller.Delete)
}
