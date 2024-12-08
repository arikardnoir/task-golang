package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"task-golang/models"
	"task-golang/services"
)

type UserController struct {
	UserService *services.UserService
}

func (ctrl *UserController) ListUsers(c echo.Context) error {
	users, err := ctrl.UserService.ListUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not fetch users"})
	}

	return c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) Register(c echo.Context) error {
	req := new(models.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Validate passwords
	if req.Password != req.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "passwords do not match"})
	}

	// Register user with address
	err := ctrl.UserService.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "user registered successfully"})
}
