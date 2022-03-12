package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/satriyoaji/echo-restful-api/controllers"
	"github.com/satriyoaji/echo-restful-api/middlewares"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello world!")
	})

	e.GET("api/employees", controllers.FetchAllEmployees, middlewares.IsAuthenticated)
	e.POST("api/employees", controllers.StoreEmployee, middlewares.IsAuthenticated)
	e.PUT("api/employees/:id", controllers.UpdateEmployee, middlewares.IsAuthenticated)
	e.DELETE("api/employees/:id", controllers.DeleteEmployee, middlewares.IsAuthenticated)

	e.GET("api/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("api/login", controllers.ActionLogin)

	return e
}
