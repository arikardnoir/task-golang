package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"task-golang/database"
	"task-golang/models"
	"task-golang/utils"
)

func RegisterUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not hash password"})
	}
	user.Password = hashedPassword

	if err := database.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create user"})
	}

	return c.JSON(http.StatusCreated, user)
}
