package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("sid")
		if err != nil {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
