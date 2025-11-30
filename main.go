package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

var version = "dev"

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		log.Println("-- / invoked --")

		return c.String(http.StatusOK, "-- Root Page --")
	})

	e.GET("/about", func(c echo.Context) error {
		log.Println("-- /about invoked --")

		return c.String(http.StatusOK, "-- About Page --")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
