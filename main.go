package main

import (
	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	component := hello("John Doe")
	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
