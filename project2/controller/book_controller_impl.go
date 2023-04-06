package controller

import (
	"net/http"
	"project-2/models/web"
	"project-2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookControllerIml struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerIml{
		BookService: bookService,
	}
}

func (controller *BookControllerIml) GetAllBook(ctx *gin.Context) {
	result, err := controller.BookService.GetAllBook()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (controller *BookControllerIml) GetBookByID(ctx *gin.Context) {
	var bookID = ctx.Param("book_id")

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	result, err := controller.BookService.GetBookByID(int64(bookIDInt))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, web.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}

func (controller *BookControllerIml) CreateBook(ctx *gin.Context) {
	var newBook web.BookRequest

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	result, err := controller.BookService.CreateBook(newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, web.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}

func (controller *BookControllerIml) UpdateBook(ctx *gin.Context) {
	var newBook web.BookRequest
	var bookID = ctx.Param("book_id")

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	result, err := controller.BookService.UpdateBook(int64(bookIDInt), newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, web.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}

func (controller *BookControllerIml) DeleteBook(ctx *gin.Context) {
	var bookID = ctx.Param("book_id")

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	result, err := controller.BookService.DeleteBook(int64(bookIDInt))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, web.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
