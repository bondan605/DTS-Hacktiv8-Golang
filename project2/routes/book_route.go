package routes

import (
	"project-2/controller"
	"project-2/repository"
	"project-2/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBookRoute(router *gin.Engine, db *gorm.DB) {
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	router.POST("/books", bookController.CreateBook)
	router.GET("/books", bookController.GetAllBook)
	router.GET("/books/:book_id", bookController.GetBookByID)
	router.PUT("/books/:book_id", bookController.UpdateBook)
	router.DELETE("/books/:book_id", bookController.DeleteBook)
}
