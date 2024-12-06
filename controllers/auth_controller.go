package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"task-golang/models"
	"task-golang/services"
)

type AuthController struct {
	AuthService *services.AuthService
}

func (ctrl *AuthController) Login(c echo.Context) error {
	req := new(models.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	token, err := ctrl.AuthService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
