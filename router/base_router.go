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

	e.GET("/redis/api", api.RedisApi().ApiIndex)
	e.POST("/redis/api/v1/record/add", api.RedisApi().Add)
	e.GET("/redis/api/v1/record/keys", api.RedisApi().GetAllKeys)
	e.GET("/redis/api/v1/record/key/:key", api.RedisApi().GetByKey)
	e.DELETE("/redis/api/v1/record/key/:key", api.RedisApi().DeleteByKey)

	e.GET("/rabbitmq/api", api.RabbitmqApi().ApiIndex)
	e.POST("/rabbitmq/api/v1/publish/:queue", api.RabbitmqApi().Publish)
	e.GET("/rabbitmq/api/v1/consume", api.RabbitmqApi().Consume)

	e.GET("/common/sync", api.CommonApi().PullFromRedisAndPublishToRabbitMQ)
}
