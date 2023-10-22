package api

import (
	v1 "github.com/kanade0404/rental-house-scraper/internal/routes/api/v1"
	"github.com/labstack/echo/v4"
)

func API(g *echo.Echo) {
	api := g.Group("/api")
	v1.V1(api.Group("/v1"))
}
