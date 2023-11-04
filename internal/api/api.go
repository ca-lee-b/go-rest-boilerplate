package api

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ca-lee-b/go-rest-boilerplate/internal/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Api struct {
	router *echo.Echo
	logger *slog.Logger

	bookHandler *handlers.BookHandler
	authHandler *handlers.AuthHandler
}

func New(handlers *handlers.Handlers, log *slog.Logger) *Api {
	e := echo.New()

	return &Api{
		router:      e,
		logger:      log,
		bookHandler: handlers.BookHandler,
		authHandler: handlers.AuthHandler,
	}
}

func (a *Api) initializeRoutes() {
	a.router.Use(middleware.Logger())

	a.router.POST("/register", a.authHandler.Register)
	a.router.POST("/login", a.authHandler.Login)

	a.router.GET("/books", a.bookHandler.GetAllBooks)
	a.router.GET("/books/:id", a.bookHandler.GetBookByIsbn)

	a.router.POST("/books", SessionMiddleware(a.bookHandler.CreateBook))
	a.router.POST("/books/:id", SessionMiddleware(a.bookHandler.UpdateBook))

	a.router.DELETE("/books/:id", SessionMiddleware(a.bookHandler.DeleteBook))
}

func (a *Api) Listen() error {
	a.initializeRoutes()

	format := fmt.Sprintf(":%v", os.Getenv("PORT"))
	err := a.router.Start(format)
	if err != nil {
		return err
	}
	return nil
}
