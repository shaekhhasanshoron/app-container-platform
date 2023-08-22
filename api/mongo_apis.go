package api

import (
	"app-container-platform/db/model"
	_type "app-container-platform/types"
	"fmt"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type MongoApiInf interface {
	ApiIndex(c echo.Context) error
	Add(c echo.Context) error
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	DeleteById(c echo.Context) error
}

type mongo_api struct{}

func MongoApi() MongoApiInf {
	return &mongo_api{}
}

func (h *mongo_api) ApiIndex(c echo.Context) error {
	response := _type.ResponseDto{
		Message: "App Container Platform Mongo Api",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *mongo_api) Add(c echo.Context) error {
	log.Println("Add")
	inputReq := new(model.RecordConfig)
	if err := c.Bind(inputReq); err != nil {
		log.Println(fmt.Errorf("[ERROR] Invalid Input!"))
		return err
	}

	err := inputReq.SaveToMongo()
	if err != nil {
		log.Println("[Error] Error Occurred while saving data to Mongo: " + err.Error())
		return c.JSON(http.StatusInternalServerError, _type.Response().Error("Error Occurred while saving data to Mongo: "+err.Error()))
	}

	response := _type.ResponseDto{
		Message: "saved successfully",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *mongo_api) GetAll(c echo.Context) error {
	log.Println("Get List")
	resp, _ := model.RecordConfig{}.GetListFromMongo(make(map[string]interface{}))

	response := _type.ResponseDto{
		Message: "",
		Data:    resp,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *mongo_api) GetById(c echo.Context) error {
	log.Println("Get")
	databaseIdStr := c.QueryParams().Get("id")
	databaseUIdStr := c.QueryParams().Get("uid")
	resp := model.RecordConfig{}
	var err error
	if databaseIdStr != "" {
		if !bson.IsObjectIdHex(databaseIdStr) {
			log.Println("Invalid Id: " + databaseIdStr)
			return c.JSON(http.StatusBadRequest, _type.Response().Error("Invalid Id: "+databaseIdStr))
		}

		resp, err = model.RecordConfig{}.GetByIdFromMongo(bson.ObjectIdHex(databaseIdStr))
		if err != nil {
			log.Println("Error occurred BY ID: " + err.Error())
			return c.JSON(http.StatusInternalServerError, _type.Response().Error("Could not found record: "+err.Error()))
		}
	} else if databaseUIdStr != "" {
		resp, err = model.RecordConfig{}.GetByUIdFromMongo(databaseUIdStr)
		if err != nil {
			log.Println("Error occurred By UID: " + err.Error())
			return c.JSON(http.StatusInternalServerError, _type.Response().Error("Could not found record: "+err.Error()))
		}
	}

	response := _type.ResponseDto{
		Message: "",
		Data:    resp,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}

func (h *mongo_api) DeleteById(c echo.Context) error {
	log.Println("Delete")
	databaseIdStr := c.Param("id")
	if databaseIdStr == "" || !bson.IsObjectIdHex(databaseIdStr) {
		log.Println("Invalid Id: " + databaseIdStr)
		return c.JSON(http.StatusBadRequest, _type.Response().Error("Invalid Id: "+databaseIdStr))
	}

	err := model.RecordConfig{}.DeleteByIdFromMongo(bson.ObjectIdHex(databaseIdStr))
	if err != nil {
		log.Println("Error occurred delete ID: " + err.Error())
		return c.JSON(http.StatusInternalServerError, _type.Response().Error("Could not delete record: "+err.Error()))
	}

	response := _type.ResponseDto{
		Message: "deleted successfully",
		Data:    nil,
	}

	return c.JSON(http.StatusOK, _type.Response().Success(response, ""))
}
