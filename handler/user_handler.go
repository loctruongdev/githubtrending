package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSignUp(c echo.Context) error {
	type User struct {
		Email    string `json: "email"`
		FullName string `json: "name"`
	}
	user := User{
		Email:    "spring@gmail.com",
		FullName: "Spring Truong",
	}
	return c.JSON(http.StatusOK, user)
}

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "Spring",
		"email": "spring@gmail.com",
	})
}
