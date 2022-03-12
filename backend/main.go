package main

import (
	"hackathon-basic-backend/config"
	"hackathon-basic-backend/database"
	"hackathon-basic-backend/routes"
)

func main() {
	database.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + config.GoDotEnvVariable("PORT")))
}
