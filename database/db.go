package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// Check the enviroment
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "test" {
		// settingds for tests (memory DB)
		DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to test database!")
		}
	} else {
		// Default Settings
		DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database!")
		}
	}
}
