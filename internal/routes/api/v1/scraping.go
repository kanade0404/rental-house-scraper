package v1

import (
	scraping2 "github.com/kanade0404/rental-house-scraper/internal/handlers/v1/scraping"
	"github.com/labstack/echo/v4"
)

func scraping(g *echo.Group) {
	g.GET("/ur", scraping2.ScrapingUR)
}
