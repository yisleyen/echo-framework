package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/string", func(c echo.Context) error {
		return c.String(http.StatusOK, "<h1>Merhaba Echo Framework</h1>")
	})
	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Merhaba Echo Framework<h1>")
	})
	e.GET("/no-content", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.Start(":8080")
}
