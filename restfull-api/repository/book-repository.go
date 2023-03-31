package repository

import (
	"github.com/fernanda-one/golang_api/entities"
	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(b entities.Book) entities.Book
	UpdateBook(b entities.Book) entities.Book
	DeleteBook(b entities.Book)
	AllBook() []entities.Book
	FindBookById(bookID uint64) entities.Book
}

type bookRepository struct {
	bookConnection *gorm.DB
}

func NewBookRepository(dbCon *gorm.DB) BookRepository {
	return &bookRepository{
		bookConnection: dbCon,
	}
}

func (db *bookRepository) InsertBook(b entities.Book) entities.Book {
	db.bookConnection.Save(&b)
	db.bookConnection.Preload("User").Find(&b)
	return b
}

func (db *bookRepository) UpdateBook(b entities.Book) entities.Book {
	db.bookConnection.Save(&b)
	db.bookConnection.Preload("User").Find(&b)
	return b
}

func (db *bookRepository) DeleteBook(b entities.Book) {
	db.bookConnection.Delete(&b)
}

func (db *bookRepository) AllBook() []entities.Book {
	var books []entities.Book
	db.bookConnection.Preload("User").Find(&books)
	return books
}

func (db *bookRepository) FindBookById(bookID uint64) entities.Book {
	var book entities.Book
	db.bookConnection.Preload("User").Find(&book, bookID)
	return book
}
