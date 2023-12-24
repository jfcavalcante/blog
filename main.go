package main

import (
	"blog/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func main() {
	e := echo.New()

	// Little bit of middlewares for housekeeping
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	// This will initiate our template renderer
	template.NewTemplateRenderer(e, "public/*.html")
	e.GET("/hello", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index", nil)
	})

	e.Logger.Fatal(e.Start(":4040"))
}
