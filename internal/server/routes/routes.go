package routes

import (
	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/services"
	"github.com/0ero-1ne/martha-server/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, jwtManager utils.JWTManager) *gin.Engine {
	router := gin.Default()
	apiRoute := router.Group("/api")

	registerTagRouter(apiRoute, services.NewTagService(db))
	registerBookRouter(apiRoute, services.NewBookService(db))
	registerAuthorRouter(apiRoute, services.NewAuthorService(db))
	registerChapterRouter(apiRoute, services.NewChapterService(db))
	registerAuthRouter(apiRoute, services.NewAuthService(db), jwtManager)
	registerUserRouter(apiRoute, services.NewUserService(db), jwtManager)

	return router
}

func registerTagRouter(globalRoute *gin.RouterGroup, service services.TagService) {
	tagRouter(globalRoute, controllers.NewTagController(service))
}

func registerBookRouter(globalRoute *gin.RouterGroup, service services.BookService) {
	bookRouter(globalRoute, controllers.NewBookController(service))
}

func registerAuthorRouter(globalRoute *gin.RouterGroup, service services.AuthorService) {
	authorRouter(globalRoute, controllers.NewAuthorController(service))
}

func registerChapterRouter(globalRoute *gin.RouterGroup, service services.ChapterService) {
	chapterRouter(globalRoute, controllers.NewChapterController(service))
}

func registerAuthRouter(globalRoute *gin.RouterGroup, service services.AuthService, jwtManager utils.JWTManager) {
	authRouter(globalRoute, controllers.NewAuthController(service, jwtManager))
}

func registerUserRouter(globalRoute *gin.RouterGroup, service services.UserService, jwtManager utils.JWTManager) {
	userRouter(globalRoute, controllers.NewUserController(service), jwtManager)
}
