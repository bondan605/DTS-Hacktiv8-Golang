package controllers

import (
	"chapter2-2/database"
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

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	RETURNING *
	`

	db := database.GetConnection()

	var bookResult Book
	result := db.QueryRow(sqlStatement, newBook.Title, newBook.Author, newBook.Desc)
	err := result.Scan(&bookResult.BookID, &bookResult.Title, &bookResult.Author, &bookResult.Desc)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResult,
	})
}

func GetBooks(ctx *gin.Context) {
	var books []Book

	sqlStatement := `
	SELECT *
	FROM books
	`

	db := database.GetConnection()

	result, err := db.Query(sqlStatement)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	for result.Next() {
		var book Book
		err := result.Scan(&book.BookID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		books = append(books, book)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
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

	sqlStatement := `
	SELECT *
	FROM books
	WHERE book_id=$1
	`

	db := database.GetConnection()

	result := db.QueryRow(sqlStatement, bookIDInt)
	err = result.Scan(&BookFounded.BookID, &BookFounded.Title, &BookFounded.Author, &BookFounded.Desc)
	if BookFounded.BookID != 0 {
		isNotFound = false
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

	sqlStatement := `
	UPDATE books
	SET title=$1, author=$2, description=$3
	WHERE book_id=$4
	RETURNING *
	`

	db := database.GetConnection()

	var bookUpdatedResult Book
	err = db.QueryRow(sqlStatement, updatedBook.Title, updatedBook.Author, updatedBook.Desc, bookIDInt).Scan(
		&bookUpdatedResult.BookID,
		&bookUpdatedResult.Title,
		&bookUpdatedResult.Author,
		&bookUpdatedResult.Desc,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if bookUpdatedResult.BookID != 0 {
		isNotFound = false
	}

	if isNotFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "not found",
			"error_message": fmt.Sprintf("book with id: %d is not found!", bookIDInt),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id: %d has been succeessfully updated", bookUpdatedResult.BookID),
		"data":    bookUpdatedResult,
	})
}

func DeleteBook(ctx *gin.Context) {
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

	sqlStatement := `
	DELETE FROM books
	WHERE book_id=$1
	RETURNING book_id
	`

	db := database.GetConnection()

	var bookIdDeleted int
	err = db.QueryRow(sqlStatement, bookIDInt).Scan(&bookIdDeleted)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if bookIdDeleted != 0 {
		isNotFound = false
	}

	if isNotFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "not found",
			"error_message": fmt.Sprintf("book with id: %d is not found!", bookIDInt),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id: %d has been succeessfully deleted", bookIdDeleted),
	})
}
