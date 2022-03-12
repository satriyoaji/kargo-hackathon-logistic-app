package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/satriyoaji/echo-restful-api/models"
	"net/http"
	"strconv"
)

func FetchAllEmployees(c echo.Context) error {
	result, err := models.FetchAllEmployees()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreEmployee(c echo.Context) error {
	name := c.FormValue("name")
	address := c.FormValue("address")
	phone := c.FormValue("phone")

	result, err := models.StoreEmployee(name, address, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateEmployee(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	name := c.FormValue("name")
	address := c.FormValue("address")
	phone := c.FormValue("phone")

	result, err := models.UpdateEmployee(int_id, name, address, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteEmployee(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteEmployee(int_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
