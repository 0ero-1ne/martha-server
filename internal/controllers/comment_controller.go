package controllers

import (
	"net/http"

	"github.com/0ero-1ne/martha-server/internal/models"
	"github.com/0ero-1ne/martha-server/internal/services"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) CommentController {
	return CommentController{
		service: service,
	}
}

func (controller CommentController) GetAll(ctx *gin.Context) {
	comments, err := controller.service.GetAll()
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (controller CommentController) GetById(ctx *gin.Context) {
	commentId := ctx.GetUint("comment_id")
	comment, err := controller.service.GetById(commentId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (controller CommentController) Create(ctx *gin.Context) {
	var comment models.Comment

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	createdComment, err := controller.service.Create(comment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, createdComment)
}

func (controller CommentController) Update(ctx *gin.Context) {
	commentId := ctx.GetUint("comment_id")
	var comment models.Comment

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	upatedComment, err := controller.service.Update(commentId, comment)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, upatedComment)
}

func (controller CommentController) Delete(ctx *gin.Context) {
	commentId := ctx.GetUint("comment_id")
	userId := ctx.GetUint("user_id")

	err := controller.service.Delete(commentId, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}

// book
func (controller CommentController) GetAllByBookId(ctx *gin.Context) {
	bookId := ctx.GetUint("book_id")

	comments, err := controller.service.GetAllByBookId(bookId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
