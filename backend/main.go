package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hackathon-basic-backend/config"
	"hackathon-basic-backend/models"
	"hackathon-basic-backend/routes"
)

func main() {
	models.Connect()

	e := routes.Init()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Logger.Fatal(e.Start(":" + config.GoDotEnvVariable("PORT")))
}
