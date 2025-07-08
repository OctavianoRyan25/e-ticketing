package controller

import (
	"net/http"

	config "github.com/OctavianoRyan25/e-ticketing/configs"
	"github.com/OctavianoRyan25/e-ticketing/models"
	"github.com/OctavianoRyan25/e-ticketing/res"
	"github.com/OctavianoRyan25/e-ticketing/util"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var req models.UserReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, res.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Data:    nil,
		})
	}

	user := new(models.User)
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, res.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid email or password",
			Data:    nil,
		})
	}

	if user.Password != req.Password {
		return c.JSON(http.StatusUnauthorized, res.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid email or password",
			Data:    nil,
		})
	}

	t, err := util.GenerateJWT(user.Role, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, res.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, res.SuccessResponse{
		Code:    http.StatusOK,
		Message: "Login successful",
		Data: map[string]string{
			"token": t,
		},
	})
}
