package main

import (
	"rest-api-mongo/config"
	"rest-api-mongo/module/database"
	"rest-api-mongo/router"

	"github.com/labstack/echo/v4"
)

func init() {
	config.Init()
	database.Connect()
}

func main() {
	envVars := config.GetEnv()
	e := echo.New()

	// Route ...
	router.Route(e)

	// Start server
	e.Logger.Fatal(e.Start(envVars.AppPort))
}
