package routes

import (
	"github.com/0ero-1ne/martha/internal/controllers"
	"github.com/0ero-1ne/martha/internal/middlewares"
	"github.com/0ero-1ne/martha/internal/services"

	"github.com/gin-gonic/gin"
)

var chapterService = services.ChapterService{}
var chapterController = controllers.NewChapterController(chapterService)

func ChapterRoutes(globalRoute *gin.RouterGroup) {
	routes := globalRoute.Group("/chapters")
	routes.GET("/", chapterController.GetAll)
	routes.GET("/:chapter_id", middlewares.ParseParamsId([]string{"chapter_id"}), chapterController.GetById)
	routes.POST("/", chapterController.Create)
	routes.PUT("/:chapter_id", middlewares.ParseParamsId([]string{"chapter_id"}), chapterController.Update)
	routes.DELETE("/:chapter_id", middlewares.ParseParamsId([]string{"chapter_id"}), chapterController.Delete)
}
