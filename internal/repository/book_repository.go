package repository

import (
	"github.com/ca-lee-b/go-rest-boilerplate/internal/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func newBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (b *BookRepository) GetAllBooks() []models.Book {
	var books []models.Book
	result := b.db.Find(&books)
	if result.Error != nil {
		return nil
	}

	return books
}

func (b *BookRepository) GetBookByIsbn(isbn string) *models.Book {
	var book *models.Book
	result := b.db.First(&book, "isbn = ?", isbn)
	if result.Error != nil {
		return nil
	}

	return book
}

func (b *BookRepository) UpdateBook(isbn string, book *models.Book) error {
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

func (b *BookRepository) CreateBook(book *models.Book) error {
	result := b.db.Create(book)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) DeleteBook(isbn string) error {
	result := b.db.Delete(&models.Book{}, isbn)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
