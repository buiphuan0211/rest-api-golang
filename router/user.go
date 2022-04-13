package router

import (
	"rest-api-mongo/controller"

	"github.com/labstack/echo/v4"
)

func user(e *echo.Echo) {
	users := e.Group("/users")
	users.GET("", controller.GetUsers)
	users.POST("", controller.CreateUser)
}
