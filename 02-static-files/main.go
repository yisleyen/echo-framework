package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()
	e.Static("/", "pages")
	e.Start(":8080")
}
