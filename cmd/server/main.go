package main

import (
	"github.com/kanade0404/rental-house-scraper/internal/config"
	"github.com/kanade0404/rental-house-scraper/internal/logger"
	"github.com/kanade0404/rental-house-scraper/internal/routes/api"
	"github.com/labstack/echo/v4"
	"os"
)

const defaultPort = "8181"

func main() {
	e := echo.New()
	e.Logger.Fatal(run(e))
}

func run(e *echo.Echo) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	_, err := config.NewConfig()
	if err != nil {
		return err
	}
	initialize(e)
	return e.Start(":" + port)
}

func initialize(e *echo.Echo) {
	api.API(e)
	logger.Init()
}
