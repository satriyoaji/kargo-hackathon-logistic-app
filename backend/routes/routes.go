package routes

import (
	"hackathon-basic-backend/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello world!")
	})

	e.GET("api/transporter/get-all-trucks", controllers.FetchAllTrucks)
	e.GET("api/transporter/get-truck/:id", controllers.FetchByIdTruck)
	e.POST("api/transporter/store-truck", controllers.StoreTruck)
	e.PUT("api/transporter/update-truck/:id", controllers.UpdateTruck)
	e.DELETE("api/transporter/delete-truck/:id", controllers.DeleteTruck)

	e.GET("api/transporter/get-all-drivers", controllers.FetchAllDrivers)
	e.GET("api/transporter/get-driver/:id", controllers.FetchByIdDriver)
	e.POST("api/transporter/store-driver", controllers.StoreDriver)
	e.PUT("api/transporter/update-driver/:id", controllers.UpdateDriver)
	e.DELETE("api/transporter/delete-driver/:id", controllers.DeleteDriver)

	e.GET("api/shipment/get-all", controllers.FetchAllShipments)
	e.POST("api/shipment/insert", controllers.AddShipment)

	e.GET("api/generate-hash/:password", controllers.GenerateHashPassword)
	// e.POST("api/login", controllers.ActionLogin)

	return e
}
