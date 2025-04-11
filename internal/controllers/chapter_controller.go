package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/0ero-1ne/martha-server/internal/models"
	"github.com/0ero-1ne/martha-server/internal/services"
)

type ChapterController struct {
	service services.ChapterService
}

func NewChapterController(service services.ChapterService) ChapterController {
	return ChapterController{
		service: service,
	}
}

func (controller ChapterController) GetAll(ctx *gin.Context) {
	chapters, err := controller.service.GetAll()
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, chapters)
}

func (controller ChapterController) GetById(ctx *gin.Context) {
	chapterId := ctx.GetUint("chapter_id")
	chapter, err := controller.service.GetById(chapterId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, chapter)
}

func (controller ChapterController) GetChaptersByBookId(ctx *gin.Context) {
	bookId := ctx.GetUint("book_id")
	chapters, err := controller.service.GetChaptersByBookId(bookId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, chapters)
}

func (controller ChapterController) Create(ctx *gin.Context) {
	var chapter models.Chapter

	if err := ctx.ShouldBindJSON(&chapter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBook, err := controller.service.Create(chapter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdBook)
}

func (controller ChapterController) Update(ctx *gin.Context) {
	chapterId := ctx.GetUint("chapter_id")
	var newChapter models.Chapter

	if err := ctx.ShouldBindJSON(&newChapter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedChapter, err := controller.service.Update(chapterId, newChapter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedChapter)
}

func (controller ChapterController) Delete(ctx *gin.Context) {
	chapterId := ctx.GetUint("chapter_id")

	if err := controller.service.Delete(chapterId); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusNoContent)
}
