package api

import (
	"fmt"
	"log/slog"

	"github.com/ca-lee-b/go-rest-boilerplate/internal/api/handlers"
	"github.com/labstack/echo/v4"
)

type Api struct {
	server *echo.Echo
	logger *slog.Logger

	bookHandler *handlers.BookHandler
}

func New(handlers *handlers.Handlers, log *slog.Logger) *Api {
	e := echo.New()

	return &Api{
		server:      e,
		logger:      log,
		bookHandler: handlers.BookHandler,
	}
}

func (a *Api) initializeRoutes() {
	a.server.GET("/books", a.bookHandler.GetAllBooks)
	a.server.GET("/books/:id", a.bookHandler.GetBookByIsbn)
	a.server.POST("/books", a.bookHandler.CreateBook)
	a.server.POST("/books/:id", a.bookHandler.UpdateBook)
}

func (a *Api) Listen(port int) error {
	a.initializeRoutes()

	format := fmt.Sprintf(":%v", port)
	err := a.server.Start(format)
	if err != nil {
		return err
	}
	return nil
}
