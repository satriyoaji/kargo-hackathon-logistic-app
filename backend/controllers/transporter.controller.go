package controllers

import (
	"github.com/labstack/echo/v4"
	"hackathon-basic-backend/models"
	"net/http"
	"strconv"
)

func FetchAllTransporters(c echo.Context) error {
	result, err := models.FetchAllTransporters()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreTransporter(c echo.Context) error {
	licenseNumber := c.FormValue("license_number")
	address := c.FormValue("address")
	phone := c.FormValue("phone")

	result, err := models.StoreTransporter(licenseNumber, address, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateTransporter(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	licenseNumber := c.FormValue("license_number")
	address := c.FormValue("address")
	phone := c.FormValue("phone")

	result, err := models.UpdateTransporter(int_id, licenseNumber, address, phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteTransporter(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteTransporter(int_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
