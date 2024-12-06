package routes

import (
	"github.com/labstack/echo/v4"
	"task-golang/controllers"
	"task-golang/middlewares"
	"task-golang/repositories"
	"task-golang/services"
)

func SetupRoutes(e *echo.Echo) {
	userRepo := &repositories.UserRepository{}
	authService := &services.AuthService{UserRepo: userRepo}
	authController := &controllers.AuthController{AuthService: authService}

	userService := &services.UserService{UserRepo: userRepo}
	userController := &controllers.UserController{UserService: userService}

	e.POST("/login", authController.Login)
	e.POST("/register", userController.Register)
	e.POST("/recover-password", authController.RecoverPassword)
	e.POST("/reset-password", authController.ResetPassword)

	// Protected route
	e.GET("/users", userController.ListUsers, middlewares.JWTMiddleware)
}
