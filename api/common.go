package api

import (
	"app-container-platform/db/cp_rabbitmq"
	"app-container-platform/db/cp_redis"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type CommonAPiInf interface {
	PullFromRedisAndPublishToRabbitMQ(c echo.Context) error
}

type commonApi struct {
}

func (c2 commonApi) PullFromRedisAndPublishToRabbitMQ(c echo.Context) error {
	keys, err := cp_redis.GetAllKeys()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error getting key from redis")
	}
	var dataList []string
	for i, key := range keys {
		data, _ := cp_redis.Get(key)
		log.Println("data from redis")
		log.Printf("%v :key: %v: %v\n", i, key, data)
		dataList = append(dataList, data)
	}
	for _, message := range dataList {
		err := cp_rabbitmq.PublishToRabbitMQ("default", message)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "error publishing in rabbitmq")
		}
	}
	return c.JSON(http.StatusOK, "operation successful")
}

func CommonApi() CommonAPiInf {
	return &commonApi{}
}
