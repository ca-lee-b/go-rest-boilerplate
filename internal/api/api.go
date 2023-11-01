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
}

func New(handlers *handlers.Handlers, log *slog.Logger) *Api {
	e := echo.New()

	return &Api{
		router:      e,
		logger:      log,
		bookHandler: handlers.BookHandler,
	}
}

func (a *Api) initializeRoutes() {
	a.router.GET("/books", a.bookHandler.GetAllBooks)
	a.router.GET("/books/:id", a.bookHandler.GetBookByIsbn)
	a.router.POST("/books", a.bookHandler.CreateBook)
	a.router.POST("/books/:id", a.bookHandler.UpdateBook)
}

func (a *Api) Listen() error {
	f, err := os.OpenFile("logs.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		a.logger.Error("Failed to open logs file")
	}
	defer f.Close()

	a.router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           f,
	}))
	a.initializeRoutes()

	format := fmt.Sprintf(":%v", os.Getenv("PORT"))
	err = a.router.Start(format)
	if err != nil {
		return err
	}
	return nil
}
