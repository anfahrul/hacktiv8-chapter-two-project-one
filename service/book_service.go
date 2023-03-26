package service

import "hacktiv8-chapter-two-project-one/models/web"

type BookService interface {
	GetAllBook() ([]web.BookResponse, error)
	GetBookByID(bookID int64) (web.BookResponse, error)
	CreateBook(book web.BookRequest) (web.BookResponse, error)
	UpdateBook(bookID int64, book web.BookRequest) (web.BookResponse, error)
	DeleteBook(bookID int64) (string, error)
}
