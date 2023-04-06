package repository

import (
	"project-2/models/entity"
)

type BookRepository interface {
	GetAllBook() ([]entity.Book, error)
	GetBookByID(bookID int64) (entity.Book, error)
	CreateBook(book entity.Book) (entity.Book, error)
	UpdateBook(bookID int64, book entity.Book) (entity.Book, error)
	DeleteBook(bookID int64) (string, error)
}
