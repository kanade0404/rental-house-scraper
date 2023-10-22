package v1

import "github.com/labstack/echo/v4"

func V1(g *echo.Group) {
	v1 := g.Group("/v1")
	scraping(v1.Group("/scraping"))
}
