package controllers

import (
	"hackathon-basic-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllShipments(c echo.Context) error {
	result, err := models.FetchAllShipments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func AddShipment(c echo.Context) error {
	shipmentNumber := c.FormValue("shipment_number")
	licenseNumber := c.FormValue("license_number")
	driverName := c.FormValue("driver_name")
	origin := c.FormValue("origin")
	destination := c.FormValue("destination")
	loadingDate := c.FormValue("loading_date")
	status := c.FormValue("status")
	action := c.FormValue("action")

	result, err := models.AddShipment(
		shipmentNumber,
		licenseNumber,
		driverName,
		origin,
		destination,
		loadingDate,
		status,
		action,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
