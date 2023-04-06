package repository

import (
	"errors"
	"fmt"
	"project-2/models/entity"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		DB: db,
	}
}

func (repository *BookRepositoryImpl) GetAllBook() ([]entity.Book, error) {
	booksResult := []entity.Book{}

	err := repository.DB.Find(&booksResult).Error
	if err != nil {
		return []entity.Book{}, err
	}

	return booksResult, nil
}

func (repository *BookRepositoryImpl) GetBookByID(bookID int64) (entity.Book, error) {
	bookResult := entity.Book{}

	err := repository.DB.First(&bookResult, "book_id = ?", bookID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, errors.New(fmt.Sprintf("Book with id %d not found", bookID))
		}

		return entity.Book{}, err
	}

	return bookResult, nil
}

func (repository *BookRepositoryImpl) CreateBook(book entity.Book) (entity.Book, error) {
	newBook := entity.Book{
		BookName: book.BookName,
		Author:   book.Author,
	}

	err := repository.DB.Create(&newBook).Error
	if err != nil {
		return entity.Book{}, err
	}

	return newBook, nil
}

func (repository *BookRepositoryImpl) UpdateBook(bookID int64, book entity.Book) (entity.Book, error) {
	bookUpdated := entity.Book{}

	err := repository.DB.First(&bookUpdated, "book_id = ?", bookID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, errors.New(fmt.Sprintf("Book with id %d not found", bookID))
		}

		return entity.Book{}, err
	}

	err = repository.DB.Model(&bookUpdated).Updates(entity.Book{BookName: book.BookName, Author: book.Author}).Error
	if err != nil {
		return entity.Book{}, err
	}

	return bookUpdated, nil
}

func (repository *BookRepositoryImpl) DeleteBook(bookID int64) (string, error) {
	bookDeleted := entity.Book{}

	err := repository.DB.First(&bookDeleted, "book_id = ?", bookID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New(fmt.Sprintf("Book with id %d not found", bookID))
		}

		return "", err
	}

	err = repository.DB.Delete(&bookDeleted, bookID).Error
	fmt.Println("err", err)
	if err != nil {
		return "", err
	}

	return "Book deleted successfully", nil
}
