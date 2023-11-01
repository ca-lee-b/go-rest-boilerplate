package handlers

import (
	"boilerplate/internal/models"
	"boilerplate/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	BookRepository *repository.BookRepo
}

func newBookHandler(repo *repository.BookRepo) *BookHandler {
	return &BookHandler{
		BookRepository: repo,
	}
}

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	books := h.BookRepository.GetAllBooks()
	if books == nil {
		return c.String(http.StatusInternalServerError, "Failed to get books")
	}

	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookByIsbn(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	book := h.BookRepository.GetBookByIsbn(id)
	if book == nil {
		return c.String(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	var book models.Book

	err := c.Bind(&book)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if book.Isbn == "" {
		return c.String(http.StatusBadRequest, "Missing Isbn Field")
	}
	if book.Author == "" {
		return c.String(http.StatusBadRequest, "Missing Author Field")
	}
	if book.Title == "" {
		return c.String(http.StatusBadRequest, "Missing Title Field")
	}

	if err := h.BookRepository.CreateBook(&book); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create book")
	}
	return c.JSON(http.StatusCreated, book)
}
