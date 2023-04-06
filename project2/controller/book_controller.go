package controller

import "github.com/gin-gonic/gin"

type BookController interface {
	GetAllBook(ctx *gin.Context)
	GetBookByID(ctx *gin.Context)
	CreateBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
}
