package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0ero-1ne/martha-server/internal/models"
	"github.com/0ero-1ne/martha-server/internal/services"
)

type TagController struct {
	service services.TagService
}

func NewTagController(service services.TagService) TagController {
	return TagController{
		service: service,
	}
}

func (controller TagController) GetAll(ctx *gin.Context) {
	tags, err := controller.service.GetAll()
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, tags)
}

func (controller TagController) GetById(ctx *gin.Context) {
	tagId := ctx.GetUint("tag_id")
	tag, err := controller.service.GetById(tagId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, tag)
}

func (controller TagController) Create(ctx *gin.Context) {
	var tag models.Tag

	if err := ctx.ShouldBindJSON(&tag); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTag, err := controller.service.Create(tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdTag)
}

func (controller TagController) Update(ctx *gin.Context) {
	tagId := ctx.GetUint("tag_id")
	var tag models.Tag

	if err := ctx.ShouldBindJSON(&tag); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTag, err := controller.service.Update(tagId, tag)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedTag)
}

func (controller TagController) Delete(ctx *gin.Context) {
	tagId := ctx.GetUint("tag_id")

	if err := controller.service.Delete(tagId); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// many2many tag:book

func (controller TagController) GetBooks(ctx *gin.Context) {
	tagId := ctx.GetUint("tag_id")
	books, err := controller.service.GetBooks(tagId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (controller TagController) AddBook(ctx *gin.Context) {
	tagId := ctx.GetUint("tag_id")
	bookId := ctx.GetUint("book_id")

	if err := controller.service.AddBook(tagId, bookId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (controller TagController) DeleteBook(ctx *gin.Context) {
	tagId := ctx.GetUint("tag_id")
	bookId := ctx.GetUint("book_id")

	if err := controller.service.DeleteBook(tagId, bookId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
