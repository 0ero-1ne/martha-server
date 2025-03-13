package controllers

import (
	"net/http"
	"server/models"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	service services.BookService
}

func NewBookController(service services.BookService) BookController {
	return BookController{
		service: service,
	}
}

func (controller BookController) GetAll(ctx *gin.Context) {
	books, err := controller.service.GetAll()

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (controller BookController) GetById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	book, err := controller.service.GetById(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (controller BookController) Create(ctx *gin.Context) {
	var book models.Book
	err := ctx.ShouldBindJSON(&book)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBook, err := controller.service.Create(book)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdBook)
}

func (controller BookController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	var book models.Book
	err = ctx.ShouldBindJSON(&book)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook, err := controller.service.Update(uint(id), book)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedBook)
}

func (controller BookController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	err = controller.service.Delete(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// many2many book:tag

func (controller BookController) GetTags(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	tags, err := controller.service.GetTags(uint(bookId))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, tags)
}

func (controller BookController) AddTag(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	tagId, err := strconv.ParseUint(ctx.Param("tag_id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'tag_id' param value"})
		return
	}

	err = controller.service.AddTag(uint(bookId), uint(tagId))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (controller BookController) DeleteTag(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	tagId, err := strconv.ParseUint(ctx.Param("tag_id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'tag_id' param value"})
		return
	}

	err = controller.service.DeleteTag(uint(bookId), uint(tagId))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// many2many book:author

func (controller BookController) GetAuthors(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	authors, err := controller.service.GetAuthors(uint(bookId))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

func (controller BookController) AddAuthor(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	authorId, err := strconv.ParseUint(ctx.Param("author_id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'author_id' param value"})
		return
	}

	err = controller.service.AddAuthor(uint(bookId), uint(authorId))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (controller BookController) DeleteAuthor(ctx *gin.Context) {
	bookId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'id' param value"})
		return
	}

	authorId, err := strconv.ParseUint(ctx.Param("author_id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'author_id' param value"})
		return
	}

	err = controller.service.DeleteAuthor(uint(bookId), uint(authorId))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
