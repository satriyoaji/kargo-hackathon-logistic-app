package main

import (
	"hackathon-basic-backend/config"
	"hackathon-basic-backend/models"
	"hackathon-basic-backend/routes"
)

func main() {
	models.Connect()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + config.GoDotEnvVariable("PORT")))
}
