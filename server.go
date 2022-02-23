package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/weather/:zip", Weather)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

type WeatherResponse struct {
	ZipCode     string `json:"ZipCode"`
	Temperature int    `json:"Temperature"`
}

func Weather(c echo.Context) error {
	zip := c.Param("zip")
	response := WeatherResponse{
		ZipCode:     zip,
		Temperature: 0,
	}
	return c.JSON(http.StatusOK, response)
}
