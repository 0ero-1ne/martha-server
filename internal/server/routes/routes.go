package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/0ero-1ne/martha-server/internal/controllers"
	"github.com/0ero-1ne/martha-server/internal/services"
	"github.com/0ero-1ne/martha-server/internal/utils"
)

func NewRouter(db *gorm.DB, jwtManager utils.JWTManager) *gin.Engine {
	router := gin.Default()
	apiRoute := router.Group("/api")

	registerTagRouter(apiRoute, services.NewTagService(db))
	registerBookRouter(apiRoute, services.NewBookService(db))
	registerAuthorRouter(apiRoute, services.NewAuthorService(db))
	registerChapterRouter(apiRoute, services.NewChapterService(db), jwtManager)
	registerAuthRouter(apiRoute, services.NewAuthService(db), jwtManager)
	registerUserRouter(apiRoute, services.NewUserService(db), jwtManager)
	registerCommentRouter(apiRoute, services.NewCommentService(db), jwtManager)
	registerCommentRatesRouter(apiRoute, services.NewCommentRateService(db), jwtManager)

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

func registerChapterRouter(
	globalRoute *gin.RouterGroup,
	service services.ChapterService,
	jwtManager utils.JWTManager,
) {
	chapterRouter(globalRoute, controllers.NewChapterController(service), jwtManager)
}

func registerAuthRouter(
	globalRoute *gin.RouterGroup,
	service services.AuthService,
	jwtManager utils.JWTManager,
) {
	authRouter(globalRoute, controllers.NewAuthController(service, jwtManager))
}

func registerUserRouter(
	globalRoute *gin.RouterGroup,
	service services.UserService,
	jwtManager utils.JWTManager,
) {
	userRouter(globalRoute, controllers.NewUserController(service), jwtManager)
}

func registerCommentRouter(
	globalRoute *gin.RouterGroup,
	service services.CommentService,
	jwtManager utils.JWTManager,
) {
	commentRouter(globalRoute, controllers.NewCommentController(service), jwtManager)
}

func registerCommentRatesRouter(
	globalRoute *gin.RouterGroup,
	service services.CommentRateService,
	jwtManager utils.JWTManager,
) {
	commentRateRouter(globalRoute, controllers.NewCommentRateController(service), jwtManager)
}
