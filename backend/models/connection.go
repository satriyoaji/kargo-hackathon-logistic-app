package models

import (
	"hackathon-basic-backend/config"
	"hackathon-basic-backend/helpers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	connection, err := gorm.Open(mysql.Open(config.GoDotEnvVariable("DB_USERNAME")+":"+config.GoDotEnvVariable("DB_PASSWORD")+"@tcp("+config.GoDotEnvVariable("DB_HOST")+":"+config.GoDotEnvVariable("DB_PORT")+")/"+config.GoDotEnvVariable("DB_NAME")+"?parseTime=True&loc=Local"), &gorm.Config{})
	helpers.OutputPanicError(err)

	DB = connection

	connection.AutoMigrate(
		Truck{},
		Driver{},
		Shipment{},
	)
}
