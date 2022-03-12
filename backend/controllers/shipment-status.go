package controllers

import (
	"hackathon-basic-backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllStatus(c echo.Context) error {
	result, err := models.GetAllStatus()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
