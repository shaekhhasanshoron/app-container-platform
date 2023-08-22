package server

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

func New() *echo.Echo {
	echoInstance := echo.New()

	mwConfig := echoprometheus.MiddlewareConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Path(), "/testurl")
		},
	}
	echoInstance.Use(echoprometheus.NewMiddlewareWithConfig(mwConfig)) // adds middleware to gather metrics

	echoInstance.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Skipping logging for health checking api
		Skipper: func(c echo.Context) bool {
			if c.Request().RequestURI == "/health" || c.Request().RequestURI == "/metrics" {
				return true
			}
			return false
		},
		Format: "[${time_rfc3339}] method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))
	echoInstance.Use(middleware.Recover())
	return echoInstance
}

func urlSkipper(c echo.Context) bool {
	if strings.HasPrefix(c.Path(), "/index") {
		return true
	} else if strings.HasPrefix(c.Path(), "/health") {
		return true
	}
	return false
}
