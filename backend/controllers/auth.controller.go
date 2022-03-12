package controllers

import (
	"hackathon-basic-backend/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

const SecretKey = "secret"

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

// func ActionLogin(c echo.Context) error {
// 	email := c.FormValue("email")
// 	password := c.FormValue("password")

// 	res, err := models.CheckLogin(email, password)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"message": err.Error(),
// 		})
// 	}
// 	if !res {
// 		return echo.ErrUnauthorized
// 	}

// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["email"] = email
// 	claims["level"] = "application"
// 	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

// 	t, err := token.SignedString([]byte(SecretKey))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"messages": err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{
// 		"message": "Login successfully !",
// 		"token":   t,
// 	})
// }
