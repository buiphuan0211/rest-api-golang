package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "get user",
	})
}
