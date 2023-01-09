package route

import (
	"fmt"
	"net/http"
	"smart-home-backend/model"
	"time"

	"smart-home-backend/database"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

/**
var allCommands []model.Command = []model.Command{{
	Uuid:     "799a817b-06ad-4217-8b40-267822e21f35",
	CreateAt: time.Now(),
	UpdateAt: time.Now(),
	Url:      "http://example.com",
}, {
	Uuid:     "799a817b-06ad-4217-8b40-267822e21f34",
	CreateAt: time.Now(),
	UpdateAt: time.Now(),
	Url:      "http://example.com",
}}
**/

func CheckCommandUuidExist(commandUuid string) bool {
	var command model.Command

	result := database.Instance.First(&command, "uuid = ?", commandUuid)

	// Exist at least 1 record, result.RowsAffected > 0, else result.RowsAffected == 0
	return result.RowsAffected != 0
}

func CreateCommand(c echo.Context) error {
	input := new(model.Command)
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

	//Insert Data into Command Database
	database.Instance.Create(&input)

	return c.JSON(http.StatusOK, model.HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Command model.Command `json:"command"`
		}{
			Command: *input,
		},
	})
}

func GetCommands(c echo.Context) error {

	//var request model.Request

	var commands []model.Command

	//Get All Data From Request Database
	database.Instance.Find(&commands)

	return c.JSON(http.StatusOK, model.HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Commands []model.Command `json:"commands"`
		}{
			Commands: commands,
		},
	})
}

func UpdateCommandByUUID(c echo.Context) error {
	input := new(model.Command)
	if err := c.Bind(input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	commandUuid := c.Param("uuid")

	if !CheckCommandUuidExist(commandUuid) {
		return c.String(http.StatusInternalServerError, "Uuid Not Exist!")
	}

	output := new(model.Command)
	database.Instance.First(&output, "uuid = ?", commandUuid)
	database.Instance.Model(&output).Where("uuid = ?", commandUuid).Updates(model.Command{UpdateAt: time.Now(), Url: input.Url})

	return c.JSON(http.StatusOK, model.HttpSuccessResponse{
		Status: "success",
		Data: struct {
			Command model.Command `json:"command"`
		}{
			Command: *output,
		},
	})
}

func DeleteCommandByUUID(c echo.Context) error {

	commandUuid := c.Param("uuid")

	if !CheckCommandUuidExist(commandUuid) {
		return c.String(http.StatusInternalServerError, "Uuid Not Exist!")
	}

	fmt.Println("Delete uuid", commandUuid)

	var command model.Command
	database.Instance.Where("uuid = ?", commandUuid).Delete(&command)

	return c.JSON(http.StatusOK, model.HttpSuccessResponse{
		Status: "success",
		Data:   nil,
	})
}
