package api

import (
	"app-container-platform/db/cp_rabbitmq"
	"app-container-platform/db/cp_redis"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type CommonAPiInf interface {
	PullFromRedisAndPublishToRabbitMQ(c echo.Context) error
}

type commonApi struct {
}

func (c2 commonApi) PullFromRedisAndPublishToRabbitMQ(c echo.Context) error {
	queue := c.QueryParam("queue")
	if queue == "" {
		queue = "default"
	}

	threads := c.QueryParam("threads")
	if threads == "" {
		threads = "1"
	}

	threadsVal, _ := strconv.Atoi(threads)

	keys, err := cp_redis.GetAllKeys()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error getting key from redis")
	}

	for i := 0; i < threadsVal; i++ {
		go syncToQueue(keys, queue)
	}

	return c.JSON(http.StatusOK, "sync has been started")
}

func syncToQueue(keys []string, queue string) {
	log.Println("Sync has been started; Queue - " + queue)

	for i, key := range keys {
		data, err := cp_redis.Get(key)
		if err != nil {
			log.Println(fmt.Sprintf("[ERROR] sync error - gettting data for key '%s': %s", key, err.Error()))
			continue
		}
		log.Printf("publishing data from redis -- %v :key: %v\n", i, key)

		err = cp_rabbitmq.PublishToRabbitMQ(queue, data)
		if err != nil {
			log.Println(fmt.Sprintf("[ERROR] sync error - publishing in rabbitmq for key '%s': %s", key, err.Error()))
			continue
		}
	}
}

func CommonApi() CommonAPiInf {
	return &commonApi{}
}
