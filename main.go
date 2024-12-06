package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"task-golang/database"
	"task-golang/database/migrations"
	"task-golang/database/seeders"
	"task-golang/routes"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

	// Start Connect
	database.ConnectDatabase()

	// Execute migrations and seeder
	migrations.Migrate(database.DB)
	seeders.SeedUsers(database.DB)

	// Set up Echo
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Set up routes
	routes.SetupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
