package seeders

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"task-golang/database"
	"task-golang/models"
)

func SeedUsers(db *gorm.DB) {
	users := []models.User{
		{
			ID:       uuid.New(),
			Name:     "Alice Doe",
			Email:    "alice@example.com",
			Password: "password123",
			Address:  "123 Main St, Bairro A, 123, Cidade X, Estado Y, 12345-678",
		},
		{
			ID:       uuid.New(),
			Name:     "Bob Smith",
			Email:    "bob@example.com",
			Password: "password456",
			Address:  "456 Elm St, Bairro B, 456, Cidade Y, Estado Z, 23456-789",
		},
		{
			ID:       uuid.New(),
			Name:     "Charlie Brown",
			Email:    "charlie@example.com",
			Password: "password789",
			Address:  "789 Oak St, Bairro C, 789, Cidade Z, Estado X, 34567-890",
		},
	}

	for _, user := range users {
		if err := database.DB.Create(&user).Error; err != nil {
			log.Printf("Failed to seed user %s: %v", user.Email, err)
		} else {
			log.Printf("Seeded user: %s", user.Email)
		}
	}
}
