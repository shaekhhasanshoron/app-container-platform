package api

import (
	"app-container-platform/db/cp_redis"
	_type "app-container-platform/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type RedisApiInf interface {
	ApiIndex(c echo.Context) error
	Add(c echo.Context) error
	GetAllKeys(c echo.Context) error
	GetByKey(c echo.Context) error
	DeleteByKey(c echo.Context) error
}

type redis_api struct{}

func RedisApi() RedisApiInf {
	return &redis_api{}
}

func (h *redis_api) ApiIndex(c echo.Context) error {
	response := _type.ResponseDto{
		Message: "App Container Platform Redis Api",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *redis_api) Add(c echo.Context) error {
	inputReq := new(_type.RedisInput)
	if err := c.Bind(inputReq); err != nil {
		log.Println(fmt.Errorf("[ERROR] Invalid Input!"))
		return err
	}

	err := cp_redis.Set(inputReq.Key, inputReq.Value)
	if err != nil {
		log.Println("[Error] Error Occurred while saving data to Redis: " + err.Error())
		return c.JSON(http.StatusInternalServerError, _type.Response().Error("Error Occurred while saving data to Redis: "+err.Error()))
	}

	response := _type.ResponseDto{
		Message: "saved successfully",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *redis_api) GetAllKeys(c echo.Context) error {
	resp, _ := cp_redis.GetAllKeys()

	response := _type.ResponseDto{
		Message: "",
		Data:    resp,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *redis_api) GetByKey(c echo.Context) error {
	key := c.Param("key")
	resp, err := cp_redis.Get(key)
	if err != nil {
		log.Println("Error occurred BY key: " + err.Error())
		return c.JSON(http.StatusInternalServerError, _type.Response().Error("Could not found record: "+err.Error()))
	}

	response := _type.ResponseDto{
		Message: "",
		Data:    resp,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *redis_api) DeleteByKey(c echo.Context) error {
	key := c.Param("key")
	err := cp_redis.Delete(key)
	if err != nil {
		log.Println("Error occurred BY key: " + err.Error())
		return c.JSON(http.StatusInternalServerError, _type.Response().Error("Could not delete record: "+err.Error()))
	}

	response := _type.ResponseDto{
		Message: "deleted successfully",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}
