package main

import (
	"github.com/BestTTTA/echo/path"
	"github.com/labstack/echo/v4"
	"net/http"
)


func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", path.GetUser)
	e.GET("/show", path.Show)
	e.POST("/save", path.Save)
	e.POST("/users", path.Users)
	e.Logger.Fatal(e.Start(":1323"))

}
