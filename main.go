package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type wordResult struct {
	Word string `json:"word"`
}

type wordList struct{}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", word)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}

func newWordList() *wordList {
	l := &wordList{}
	return l
}

func (l *wordList) Get() *wordResult {
	return &wordResult{Word: "dog"}
}

// Handler
func word(c echo.Context) error {
	list := newWordList()
	word := list.Get()
	return c.JSON(http.StatusOK, word)
}
