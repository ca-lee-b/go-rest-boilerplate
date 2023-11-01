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

func (b *BookRepo) CreateBook(book *models.Book) error {
	result := b.db.Create(book)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
