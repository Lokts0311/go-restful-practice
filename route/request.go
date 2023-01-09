package route

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"smart-home-backend/database"
	"smart-home-backend/model"
)

/*
var allRequests []model.Request = []model.Request{{
	Uuid:     "b6615256-c658-4ee8-afa8-7b4647a460f8",
	CreateAt: time.Now(),
	UpdateAt: time.Now(),
	Name:     "Request 1",
}, {
	Uuid:     "d4723050-2c37-4fde-8000-af74b6712075",
	CreateAt: time.Now(),
	UpdateAt: time.Now(),
	Name:     "Request 2",
}}*/

func CheckRequestUuidExist(requestUuid string) bool {
	var requests model.Request

	result := database.Instance.First(&requests, "uuid = ?", requestUuid)

	// Exist at least 1 record, result.RowsAffected > 0, else result.RowsAffected == 0
	return result.RowsAffected != 0
}

func CreateRequest(c echo.Context) error {
	input := new(model.Request)
	if err := c.Bind(input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	_uuid, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		return c.String(http.StatusInternalServerError, "Fail to generate uuid")
	}
	input.Uuid = _uuid.String()
	input.CreateAt = time.Now()
	input.UpdateAt = time.Now()

	//Insert Data into Request Database
	database.Instance.Create(&input)

	return c.JSON(http.StatusOK, model.HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Request model.Request `json:"request"`
		}{
			Request: *input,
		},
	})
}

func GetRequests(c echo.Context) error {

	//var request model.Request

	var requests []model.Request

	//Get All Data From Request Database
	database.Instance.Find(&requests)

	return c.JSON(http.StatusOK, model.HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Requests []model.Request `json:"request"`
		}{
			Requests: requests,
		},
	})
}

func UpdateRequestByUUID(c echo.Context) error {
	input := new(model.Request)
	if err := c.Bind(input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	requestUuid := c.Param("uuid")

	if !CheckRequestUuidExist(requestUuid) {
		return c.String(http.StatusInternalServerError, "Uuid Not Exist!")
	}

	output := new(model.Request)
	database.Instance.First(&output, "uuid = ?", requestUuid)
	database.Instance.Model(&output).Where("uuid = ?", requestUuid).Updates(model.Request{UpdateAt: time.Now(), Name: input.Name})

	return c.JSON(http.StatusOK, model.HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Request model.Request `json:"request"`
		}{
			Request: *output,
		},
	})
}

func DeleteRequestByUUID(c echo.Context) error {
	requestUuid := c.Param("uuid")

	if !CheckRequestUuidExist(requestUuid) {
		return c.String(http.StatusInternalServerError, "Uuid Not Exist!")
	}

	fmt.Println("Delete uuid", requestUuid)

	var request model.Request
	database.Instance.Where("uuid = ?", requestUuid).Delete(&request)

	return c.JSON(http.StatusOK, model.HttpSuccessResponse{
		Status: "success",
		Data:   nil,
	})
}
