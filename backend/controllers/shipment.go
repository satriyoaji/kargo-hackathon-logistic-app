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
