package controllers

import (
	"net/http"

	"github.com/0ero-1ne/martha-server/internal/models"
	"github.com/0ero-1ne/martha-server/internal/services"
	"github.com/gin-gonic/gin"
)

type CommentRateController struct {
	service services.CommentRateService
}

func NewCommentRateController(service services.CommentRateService) CommentRateController {
	return CommentRateController{
		service: service,
	}
}

func (controller CommentRateController) GetAll(ctx *gin.Context) {
	commentRates, err := controller.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, commentRates)
}

func (controller CommentRateController) Create(ctx *gin.Context) {
	var commentRate models.CommentsRates
	if err := ctx.ShouldBindJSON(&commentRate); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if commentRate.UserId != ctx.GetUint("user_id") {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	createdCommentRate, err := controller.service.Create(commentRate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, createdCommentRate)
}

func (controller CommentRateController) Update(ctx *gin.Context) {
	var commentRate models.CommentsRates
	if err := ctx.ShouldBindJSON(&commentRate); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if commentRate.UserId != ctx.GetUint("user_id") {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	upatedCommentRate, err := controller.service.Update(commentRate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, upatedCommentRate)
}

func (controller CommentRateController) Delete(ctx *gin.Context) {
	commentId := ctx.GetUint("comment_id")
	userId := ctx.GetUint("comment_user_id")

	if userId != ctx.GetUint("user_id") {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := controller.service.Delete(commentId, userId); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
