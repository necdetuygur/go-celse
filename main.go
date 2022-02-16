package main

import (
	"celsetest/celse"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		celse.Set(c.Response(), c.Request(), "k1", "v1")
		return c.NoContent(http.StatusOK)
	})

	e.GET("/get", func(c echo.Context) error {
		k1 := celse.Get(c.Response(), c.Request(), "k1")
		return c.String(http.StatusOK, k1)
	})

	e.Logger.Fatal(e.Start("127.0.0.1:1995"))
}
