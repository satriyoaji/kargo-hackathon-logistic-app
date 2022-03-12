package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satriyoaji/echo-restful-api/config"
	"github.com/satriyoaji/echo-restful-api/helpers"
)

var DB *sql.DB

func Init() {

	connectionString := config.GoDotEnvVariable("DB_USERNAME") +
		":" + config.GoDotEnvVariable("DB_PASSWORD") +
		"@tcp(" + config.GoDotEnvVariable("DB_HOST") +
		":" + config.GoDotEnvVariable("DB_PORT") +
		")/" + config.GoDotEnvVariable("DB_NAME") +
		"?parseTime=True&loc=Local"
	connection, err := sql.Open("mysql", connectionString)
	helpers.OutputPanicError(err)

	DB = connection

	//connection.AutoMigrate('...')
}

func GetConnection() *sql.DB {
	return DB
}
