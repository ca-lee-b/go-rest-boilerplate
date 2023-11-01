package handlers

import (
	"boilerplate/internal/repository"
	"log/slog"
)

type Handlers struct {
	BookHandler *BookHandler

	Logger *slog.Logger
}

func New(repository *repository.Repo, logger *slog.Logger) *Handlers {
	return &Handlers{
		BookHandler: newBookHandler(&repository.BookRepo),
		Logger:      logger,
	}
}
