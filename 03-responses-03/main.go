package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/", "public")
	e.GET("/gopher/:name", func(c echo.Context) error {
		prm := c.Param("name")
		if prm == "svg" {
			return c.Inline("images/gopher.svg", "gopher.svg")
		}
		if prm == "jpg" {
			return c.File("images/gopher.jpg")
		}
		if prm == "att" {
			return c.Attachment("images/gopher.svg", "gopher.svg")
		}
		return c.HTML(http.StatusNotFound, "<p>Hatalı parametre</p>")
	})
	e.Start(":8080")
}
