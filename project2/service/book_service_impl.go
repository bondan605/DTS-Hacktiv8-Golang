package service

import (
	"project-2/models/entity"
	"project-2/models/web"
	"project-2/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}

func (service *BookServiceImpl) GetAllBook() ([]web.BookResponse, error) {
	booksResponse := []web.BookResponse{}

	result, err := service.BookRepository.GetAllBook()
	if err != nil {
		return []web.BookResponse{}, err
	}

	for _, val := range result {
		booksResponse = append(booksResponse, web.BookResponse(val))
	}

	return booksResponse, nil
}

func (service *BookServiceImpl) GetBookByID(bookID int64) (web.BookResponse, error) {
	result, err := service.BookRepository.GetBookByID(bookID)
	if err != nil {
		return web.BookResponse{}, err
	}

	return web.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (service *BookServiceImpl) CreateBook(book web.BookRequest) (web.BookResponse, error) {
	newBook := entity.Book{
		BookName: book.BookName,
		Author:   book.Author,
	}

	result, err := service.BookRepository.CreateBook(newBook)
	if err != nil {
		return web.BookResponse{}, err
	}

	return web.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (service *BookServiceImpl) UpdateBook(bookID int64, book web.BookRequest) (web.BookResponse, error) {
	newBook := entity.Book{
		BookName: book.BookName,
		Author:   book.Author,
	}

	result, err := service.BookRepository.UpdateBook(bookID, newBook)
	if err != nil {
		return web.BookResponse{}, err
	}

	return web.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (service *BookServiceImpl) DeleteBook(bookID int64) (string, error) {
	result, err := service.BookRepository.DeleteBook(bookID)
	if err != nil {
		return "", err
	}

	return result, nil
}
