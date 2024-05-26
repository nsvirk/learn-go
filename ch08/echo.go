package ch08

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func LearnEcho() {
	e := echo.New()

	e.GET("/", rootHandler)
	e.GET("/query", queryHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func rootHandler(c echo.Context) error {
	return c.JSON(200, map[string]string{"message": "Hello, World!"})
}

func queryHandler(c echo.Context) error {
	name := c.QueryParam("name")
	return c.JSON(200, map[string]string{"message": fmt.Sprintf("Hello, %s!", name)})
}
