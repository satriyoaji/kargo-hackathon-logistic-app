package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/satriyoaji/echo-restful-api/controllers"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(controllers.SecretKey),
})
