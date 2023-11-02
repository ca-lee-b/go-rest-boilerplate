package handlers

import (
	"log/slog"

	"github.com/ca-lee-b/go-rest-boilerplate/internal/repository"
)

type Handlers struct {
	BookHandler *BookHandler
	AuthHandler *AuthHandler

	Logger *slog.Logger
}

func New(repository *repository.Repo, logger *slog.Logger) *Handlers {
	return &Handlers{
		BookHandler: newBookHandler(repository.BookRepo),
		AuthHandler: newAuthHandler(repository.UserRepo, repository.SessionRepo),
		Logger:      logger,
	}
}
