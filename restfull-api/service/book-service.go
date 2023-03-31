package service

import (
	"fmt"
	"github.com/fernanda-one/golang_api/Dto"
	"github.com/fernanda-one/golang_api/entities"
	"github.com/fernanda-one/golang_api/repository"
	"github.com/mashingan/smapping"
	"log"
)

type BookService interface {
	Insert(b Dto.BookCreateDto) entities.Book
	Update(b Dto.BookUpdateDto) entities.Book
	Delete(b entities.Book)
	All() []entities.Book
	FindById(bookID uint64) entities.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (service *bookService) Insert(b Dto.BookCreateDto) entities.Book {
	book := entities.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.bookRepository.InsertBook(book)
	return res
}

func (service *bookService) Update(b Dto.BookUpdateDto) entities.Book {
	book := entities.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.bookRepository.UpdateBook(book)
	return res
}

func (service *bookService) Delete(b entities.Book) {
	service.bookRepository.DeleteBook(b)
}

func (service *bookService) All() []entities.Book {
	return service.bookRepository.AllBook()
}
func (service *bookService) FindById(bookID uint64) entities.Book {
	return service.bookRepository.FindBookById(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	book := service.bookRepository.FindBookById(bookID)
	id := fmt.Sprintf("%v", book.UserID)
	return userID == id
}
