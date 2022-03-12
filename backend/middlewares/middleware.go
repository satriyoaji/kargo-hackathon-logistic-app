package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
	"hackathon-basic-backend/controllers"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(controllers.SecretKey),
})
