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

	e.GET("api/transporters", controllers.FetchAllTransporters)
	e.POST("api/transporters", controllers.StoreTransporter)
	e.PUT("api/transporters/:id", controllers.UpdateTransporter)
	e.DELETE("api/transporters/:id", controllers.DeleteTransporter)

	e.GET("api/shipment/get-all", controllers.FetchAllShipments, middlewares.IsAuthenticated)

	e.GET("api/generate-hash/:password", controllers.GenerateHashPassword)
	// e.POST("api/login", controllers.ActionLogin)

	return e
}
