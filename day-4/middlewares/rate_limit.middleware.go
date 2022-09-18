package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"os"
	"strconv"
)

func RateLimit(e *echo.Echo) {
	var limit float64
	var err error
	if limit, err = strconv.ParseFloat(os.Getenv("HTTP_RATE_LIMIT"), 64); err != nil {
		limit = 5000
	}
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(limit))))
}
