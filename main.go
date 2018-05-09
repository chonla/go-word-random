package main

import (
	"bufio"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type wordResult struct {
	Word string `json:"word"`
}

type wordList struct {
	words []string
}

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
	rand.Seed(time.Now().Unix())
	words, _ := readLines("./words.txt")

	l := &wordList{
		words: words,
	}

	return l
}

func (l *wordList) Get() *wordResult {
	w := l.words[rand.Intn(len(l.words))]
	return &wordResult{Word: w}
}

// Handler
func word(c echo.Context) error {
	list := newWordList()
	word := list.Get()
	return c.JSON(http.StatusOK, word)
}

// utils
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
