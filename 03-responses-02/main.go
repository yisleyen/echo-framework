package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	plist := make([]Person, 0)
	p := Person{
		Name:    "Yusuf",
		Surname: "İşleyen",
		Age:     30,
	}
	plist = append(plist, p)
	plist = append(plist, p)
	plist = append(plist, p)
	plist = append(plist, p)
	plist = append(plist, p)
	e := echo.New()
	e.GET("/json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, plist)
	})
	e.GET("/xml", func(c echo.Context) error {
		return c.XML(http.StatusOK, plist)
	})
	e.Start(":8080")
}

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}
