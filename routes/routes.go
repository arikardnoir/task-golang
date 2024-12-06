package routes

import (
	"github.com/labstack/echo/v4"
	"task-golang/controllers"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/register", controllers.RegisterUser)
}
