package main

import (
	"app-container-platform/config"
	"app-container-platform/db/cp_mongodb"
	"app-container-platform/db/cp_rabbitmq"
	"app-container-platform/db/cp_redis"
	"app-container-platform/router"
	"app-container-platform/server"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	config.InitEnvironmentVariables()
	if config.ConnectMongo == "true" {
		_ = cp_mongodb.InitMongoDbWriteConnection()
		_ = cp_mongodb.InitMongoDbReadConnection()
		cp_mongodb.InitDBCollections()
	}

	if config.ConnectRedis == "true" {
		if config.RedisConnectionType == "SENTINEL" {
			_ = cp_redis.InitRedisSentinelConnection()
		} else {
			_ = cp_redis.InitRedisWriteConnection()
			_ = cp_redis.InitRedisReadConnection()
		}
	}

	if config.ConnectRabbitMQ == "true" {
		_ = cp_rabbitmq.InitRabbitMQConnection()
	}

	srv := server.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}
	srv.Renderer = renderer
	router.Routes(srv)
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
