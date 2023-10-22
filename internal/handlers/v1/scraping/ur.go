package scraping

import "github.com/labstack/echo/v4"

func ScrapingUR(c echo.Context) error {
	return c.String(200, "ok")
}
