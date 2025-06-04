package controllers

import (
	"net/http"

	"github.com/0ero-1ne/martha-server/internal/models"
	"github.com/0ero-1ne/martha-server/internal/services"
	"github.com/gin-gonic/gin"
)

type BookRateController struct {
	service services.BookRateService
}

func NewBookRateController(service services.BookRateService) BookRateController {
	return BookRateController{
		service: service,
	}
}

func (controller BookRateController) GetAll(ctx *gin.Context) {
	bookRates, err := controller.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, bookRates)
}

func (controller BookRateController) Create(ctx *gin.Context) {
	var bookRate models.BooksRates
	if err := ctx.ShouldBindJSON(&bookRate); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if bookRate.UserId != ctx.GetUint("user_id") {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	createdBookRate, err := controller.service.Create(bookRate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, createdBookRate)
}

func (controller BookRateController) Update(ctx *gin.Context) {
	var bookRate models.BooksRates
	if err := ctx.ShouldBindJSON(&bookRate); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if bookRate.UserId != ctx.GetUint("user_id") {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	upatedBookRate, err := controller.service.Update(bookRate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, upatedBookRate)
}

func (controller BookRateController) Delete(ctx *gin.Context) {
	bookId := ctx.GetUint("book_id")
	userId := ctx.GetUint("book_user_id")

	if userId != ctx.GetUint("user_id") {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := controller.service.Delete(bookId, userId); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
