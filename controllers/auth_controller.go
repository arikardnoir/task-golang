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

type RecoveryRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
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

func (ctrl *AuthController) RecoverPassword(c echo.Context) error {
	req := new(RecoveryRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	token, err := ctrl.AuthService.GenerateRecoveryToken(req.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "recovery token generated", "token": token})
}

func (ctrl *AuthController) ResetPassword(c echo.Context) error {
	req := new(ResetPasswordRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	err := ctrl.AuthService.ResetPassword(req.Token, req.NewPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "password reset successfully"})
}
