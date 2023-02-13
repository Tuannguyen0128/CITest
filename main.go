package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	School string `json:"school"`
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//make a json of Person
	json, err := json.Marshal(&Person{
		Name:   "Nguyen Van B",
		Age:    18,
		School: "Hong Bang university",
	})
	if err != nil {
		return  
	}

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, string(json))
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
