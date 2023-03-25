package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusOK, gin.H{
		"data": newBook,
	})
}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": BookDatas,
	})
}

func GetBookByID(ctx *gin.Context) {
	var BookFounded Book
	var bookID = ctx.Param("book_id")
	var isNotFound = true

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":   "Book ID must be a number",
			"error_messages": err,
		})
		return
	}

	for _, book := range BookDatas {
		if bookIDInt == book.BookID {
			isNotFound = false
			BookFounded = book
			break
		}
	}

	if isNotFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "not found",
			"error_message": fmt.Sprintf("book with id: %d is not found!", bookIDInt),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": BookFounded,
	})
}

func UpdateBook(ctx *gin.Context) {
	var updatedBook Book
	var bookID = ctx.Param("book_id")
	var isNotFound = true

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":   "Book ID must be a number",
			"error_messages": err,
		})
		return
	}

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookIDInt == book.BookID {
			isNotFound = false
			BookDatas[i] = updatedBook
			BookDatas[i].BookID = bookIDInt
			break
		}
	}

	if isNotFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "not found",
			"error_message": fmt.Sprintf("book with id: %d is not found!", bookIDInt),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id: %d has been succeessfully updated", bookIDInt),
	})
}

func DeleteBook(ctx *gin.Context) {
	var bookID = ctx.Param("book_id")
	var isNotFound = true
	var bookIndex int

	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":   "Book ID must be a number",
			"error_messages": err,
		})
		return
	}

	for i, book := range BookDatas {
		if bookIDInt == book.BookID {
			isNotFound = false
			bookIndex = i
			break
		}
	}

	if isNotFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "not found",
			"error_message": fmt.Sprintf("book with id: %d is not found!", bookIDInt),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id: %d has been succeessfully deleted", bookIDInt),
	})
}
