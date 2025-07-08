package routes

import (
	"github.com/OctavianoRyan25/e-ticketing/configs"
	"github.com/OctavianoRyan25/e-ticketing/controller"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.POST("/login", controller.Login)

	api := e.Group("/api")
	api.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(configs.JWT_SECRET),
	}))
	api.POST("/terminal", controller.CreateTerminal)

}
