package main //import github.com/airylinus/echo-rater

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	type PathLimiter struct {
		Path string
		Max  float64
	}

	conf := []PathLimiter{
		{
			Path: "/test/max1",
			Max:  float64(10),
		},
		{
			Path: "/test/max2",
			Max:  float64(2),
		},
	}
	for _, c := range conf {
		e.GET(c.Path, TestHandler, GatewayLimitMiddleware(c.Max))
	}

	e.Start(":4444")
}

func TestHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GatewayLimitMiddleware(max float64) echo.MiddlewareFunc {
	var tbOptions limiter.ExpirableOptions
	tbOptions.DefaultExpirationTTL = time.Second
	tbOptions.ExpireJobInterval = 0
	lmt := tollbooth.NewLimiter(max, &tbOptions)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			httpError := tollbooth.LimitByRequest(lmt, c.Response(), c.Request())
			if httpError != nil {
				// @todo
				return c.String(httpError.StatusCode, httpError.Message)
			}
			return next(c)
		})
	}
}
