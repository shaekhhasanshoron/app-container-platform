package api

import (
	_type "app-container-platform/types"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HomeInf interface {
	Index(c echo.Context) error
	ApiIndex(c echo.Context) error
	Health(c echo.Context) error
}

type home struct{}

func GeneralApi() HomeInf {
	return &home{}
}

func (h *home) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"appName": "App Container Platform",
	})
}

func (h *home) ApiIndex(c echo.Context) error {
	response := _type.ResponseDto{
		Message: "App Container Platform Api",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *home) Health(c echo.Context) error {
	return c.String(http.StatusOK, "App Container Platform I am live!")
}
