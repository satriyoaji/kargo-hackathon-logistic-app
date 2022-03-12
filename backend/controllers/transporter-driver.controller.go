package controllers

import (
	"github.com/labstack/echo/v4"
	"hackathon-basic-backend/models"
	"net/http"
	"strconv"
)

func FetchAllDrivers(c echo.Context) error {
	result, err := models.FetchAllDrivers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchByIdDriver(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, err := models.FetchByIdDriver(int_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreDriver(c echo.Context) error {
	name := c.FormValue("name")
	phone := c.FormValue("phone")
	idCard := c.FormValue("id_card")
	license := c.FormValue("license")

	result, err := models.StoreDriver(name, phone, idCard, license)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDriver(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	name := c.FormValue("name")
	phone := c.FormValue("phone")
	idCard := c.FormValue("id_card")
	license := c.FormValue("license")
	status := c.FormValue("status")

	result, err := models.Response{}, nil
	if status == "Active" || status == "active" {
		result, err = models.UpdateDriver(int_id, name, phone, idCard, license, true)
	} else {
		result, err = models.UpdateDriver(int_id, name, phone, idCard, license, false)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteDriver(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteDriver(int_id)
	if err != nil {
		if result.Message != "" {
			return c.JSON(http.StatusInternalServerError, result)
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
