package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/0ero-1ne/martha-server/internal/models"
	"github.com/0ero-1ne/martha-server/internal/services"
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
	params := models.BookUrlParams{
		WithTags:      len(ctx.Query("withTags")) != 0,
		WithAuthors:   len(ctx.Query("withAuthors")) != 0,
		WithComments:  len(ctx.Query("withComments")) != 0,
		WithChapters:  len(ctx.Query("withChapters")) != 0,
		WithBookRates: len(ctx.Query("withBookRates")) != 0,
		Query:         strings.TrimSpace(ctx.Query("query")),
		Tags:          strings.TrimSpace(ctx.Query("tags")),
	}
	books, err := controller.service.GetAll(params)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (controller BookController) GetById(ctx *gin.Context) {
	bookId := ctx.GetUint("book_id")

	params := models.BookUrlParams{
		WithTags:      len(ctx.Query("withTags")) != 0,
		WithAuthors:   len(ctx.Query("withAuthors")) != 0,
		WithComments:  len(ctx.Query("withComments")) != 0,
		WithChapters:  len(ctx.Query("withChapters")) != 0,
		WithBookRates: len(ctx.Query("withBookRates")) != 0,
	}

	book, err := controller.service.GetById(bookId, params)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (controller BookController) Create(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
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
	bookId := ctx.GetUint("book_id")
	var newBook models.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook, err := controller.service.Update(bookId, newBook)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedBook)
}

func (controller BookController) Delete(ctx *gin.Context) {
	bookId := ctx.GetUint("book_id")

	if err := controller.service.Delete(bookId); err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// many2many book:tag

func (controller BookController) GetTags(ctx *gin.Context) {
	bookId := ctx.GetUint("book_id")
	tags, err := controller.service.GetTags(bookId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, tags)
}

func (controller BookController) AddTag(ctx *gin.Context) {
	bookId := ctx.GetUint("book_id")
	tagId := ctx.GetUint("tag_id")

	if err := controller.service.AddTag(bookId, tagId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (controller BookController) DeleteTag(ctx *gin.Context) {
	bookId := ctx.GetUint("book_id")
	tagId := ctx.GetUint("tag_id")

	if err := controller.service.DeleteTag(bookId, tagId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// many2many book:author

func (controller BookController) GetAuthors(ctx *gin.Context) {
	bookId := ctx.GetUint("book_id")
	authors, err := controller.service.GetAuthors(bookId)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, authors)
}

func (controller BookController) AddAuthor(ctx *gin.Context) {
	bookId := ctx.GetUint("book_id")
	authorId := ctx.GetUint("author_id")

	if err := controller.service.AddAuthor(bookId, authorId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (controller BookController) DeleteAuthor(ctx *gin.Context) {
	bookId := ctx.GetUint("book_id")
	authorId := ctx.GetUint("author_id")

	if err := controller.service.DeleteAuthor(bookId, authorId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
