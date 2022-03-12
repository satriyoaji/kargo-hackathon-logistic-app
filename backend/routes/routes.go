package routes

import (
	"hackathon-basic-backend/controllers"
	"hackathon-basic-backend/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello world!")
	})

	e.GET("api/transporters", controllers.FetchAllTransporters, middlewares.IsAuthenticated)
	e.POST("api/transporters", controllers.StoreTransporter, middlewares.IsAuthenticated)
	e.PUT("api/transporters/:id", controllers.UpdateTransporter, middlewares.IsAuthenticated)
	e.DELETE("api/transporters/:id", controllers.DeleteTransporter, middlewares.IsAuthenticated)

	e.GET("api/generate-hash/:password", controllers.GenerateHashPassword)
	// e.POST("api/login", controllers.ActionLogin)

	return e
}
