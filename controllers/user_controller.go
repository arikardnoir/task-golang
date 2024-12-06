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

func (ctrl *UserController) Register(c echo.Context) error {
	req := new(models.User)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	err := ctrl.UserService.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "user registered successfully"})
}
