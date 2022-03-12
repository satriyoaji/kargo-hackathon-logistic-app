package controllers

import (
	"github.com/labstack/echo/v4"
	"hackathon-basic-backend/models"
	"net/http"
	"strconv"
)

func FetchAllTrucks(c echo.Context) error {
	result, err := models.FetchAllTrucks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchByIdTruck(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, err := models.FetchByIdTruck(int_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreTruck(c echo.Context) error {
	licenseNumber := c.FormValue("license_number")
	truckType := c.FormValue("truck_type")
	plateType := c.FormValue("plate_type")
	productionYear := c.FormValue("production_year")

	result, err := models.StoreTruck(licenseNumber, truckType, plateType, productionYear)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateTruck(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	licenseNumber := c.FormValue("license_number")
	truckType := c.FormValue("truck_type")
	plateType := c.FormValue("plate_type")
	productionYear := c.FormValue("production_year")
	status := c.FormValue("status")

	result, err := models.Response{}, nil
	if status == "Active" || status == "active" {
		result, err = models.UpdateTruck(int_id, licenseNumber, truckType, plateType, productionYear, true)
	} else {
		result, err = models.UpdateTruck(int_id, licenseNumber, truckType, plateType, productionYear, false)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteTruck(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteTruck(int_id)
	if err != nil {
		if result.Message != "" {
			return c.JSON(http.StatusInternalServerError, result)
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
