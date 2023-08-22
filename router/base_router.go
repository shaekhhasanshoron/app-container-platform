package router

import (
	"app-container-platform/api"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/", api.GeneralApi().Index)
	e.GET("/api", api.GeneralApi().ApiIndex)
	e.GET("/health", api.GeneralApi().Health)

	e.GET("/mongo/api", api.MongoApi().ApiIndex)
	e.POST("/mongo/api/v1/record/add", api.MongoApi().Add)
	e.GET("/mongo/api/v1/record/list", api.MongoApi().GetAll)
	e.GET("/mongo/api/v1/record/get", api.MongoApi().GetById)
	e.DELETE("/mongo/api/v1/record/delete/:id", api.MongoApi().DeleteById)
}
