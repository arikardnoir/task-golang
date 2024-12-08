package tests

import (
	"task-golang/database"
	"task-golang/database/migrations"
)

func SetupTestDB() {
	database.ConnectDatabase()      // Connect to a databse
	migrations.Migrate(database.DB) // Run the migrations
}
