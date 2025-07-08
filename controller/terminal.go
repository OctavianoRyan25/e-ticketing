package controller

import (
	"net/http"

	config "github.com/OctavianoRyan25/e-ticketing/configs"
	"github.com/OctavianoRyan25/e-ticketing/models"
	"github.com/OctavianoRyan25/e-ticketing/res"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateTerminal(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	role := claims["role"].(string)
	if role != "admin" {
		return c.JSON(http.StatusForbidden, res.ErrorResponse{
			Code:    http.StatusForbidden,
			Message: "Access denied",
			Data:    nil,
		})
	}

	var req models.TerminalReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, res.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Data:    nil,
		})
	}

	terminal := new(models.Terminal)

	terminal.Name = req.Name
	terminal.Location = req.Location

	if err := config.DB.Create(&terminal).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, res.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create terminal",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, res.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Terminal created successfully",
		Data:    terminal,
	})
}
