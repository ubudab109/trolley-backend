package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, this is echo!")
	})

	return e
}
