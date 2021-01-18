package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/main", func(c echo.Context) error {
		return c.String(http.StatusOK, "Merhaba Echo Framework")
	})
	e.Start(":8080")
}
