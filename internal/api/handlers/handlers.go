package handlers

import (
	"log/slog"

	"github.com/ca-lee-b/go-rest-boilerplate/internal/repository"
)

type Handlers struct {
	BookHandler *BookHandler

	Logger *slog.Logger
}

func New(repository *repository.Repo, logger *slog.Logger) *Handlers {
	return &Handlers{
		BookHandler: newBookHandler(&repository.BookRepository),
		Logger:      logger,
	}
}
