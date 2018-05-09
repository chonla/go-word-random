package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type wordResult struct {
	Word string `json:"word"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", word)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func word(c echo.Context) error {
	word := &wordResult{Word: "cat"}
	return c.JSON(http.StatusOK, word)
}
