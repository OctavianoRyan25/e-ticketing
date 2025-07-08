package main

import (
	"github.com/OctavianoRyan25/e-ticketing/configs"
	"github.com/OctavianoRyan25/e-ticketing/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	configs.ConnectDB()
	configs.AutoMigrate()

	e := echo.New()
	routes.InitRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
