package repository

import (
	"github.com/ca-lee-b/go-rest-boilerplate/internal/models"

	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func newBookRepository(db *gorm.DB) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (b *BookRepo) GetAllBooks() []models.Book {
	var books []models.Book
	result := b.db.Find(&books)
	if result.Error != nil {
		return nil
	}

	return books
}

func (b *BookRepo) GetBookByIsbn(isbn string) *models.Book {
	var book *models.Book
	result := b.db.First(&book, "isbn = ?", isbn)
	if result.Error != nil {
		return nil
	}

	return book
}

func (b *BookRepo) UpdateBook(isbn string, book *models.Book) error {
	result := b.db.Model(&models.Book{}).Where("isbn = ?", isbn).Updates(models.Book{
		Isbn:   book.Isbn,
		Title:  book.Title,
		Author: book.Author,
		Price:  book.Price,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepo) CreateBook(book *models.Book) error {
	result := b.db.Create(book)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
