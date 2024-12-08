package seeders

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"task-golang/models"
)

func SeedUsers(db *gorm.DB) {
	users := []struct {
		Name     string
		Email    string
		Password string
		Address  models.Address
	}{
		{
			Name:     "Alice Doe",
			Email:    "alice@example.com",
			Password: "password123",
			Address: models.Address{
				Street:       "123 Main St",
				Neighborhood: "Bairro A",
				Number:       "123",
				City:         "Cidade X",
				State:        "Estado Y",
				CEP:          "12345-678",
			},
		},
		{
			Name:     "Bob Smith",
			Email:    "bob@example.com",
			Password: "password456",
			Address: models.Address{
				Street:       "456 Elm St",
				Neighborhood: "Bairro B",
				Number:       "456",
				City:         "Cidade Y",
				State:        "Estado Z",
				CEP:          "23456-789",
			},
		},
		{
			Name:     "Charlie Brown",
			Email:    "charlie@example.com",
			Password: "password789",
			Address: models.Address{
				Street:       "789 Oak St",
				Neighborhood: "Bairro C",
				Number:       "789",
				City:         "Cidade Z",
				State:        "Estado X",
				CEP:          "34567-890",
			},
		},
	}

	for _, u := range users {
		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Failed to hash password for user %s: %v", u.Email, err)
			continue
		}

		// Create user entity
		user := models.User{
			ID:       uuid.New(),
			Name:     u.Name,
			Email:    u.Email,
			Password: string(hashedPassword),
		}

		// Begin transaction
		tx := db.Begin()

		// Save user
		if err := tx.Create(&user).Error; err != nil {
			log.Printf("Failed to seed user %s: %v", u.Email, err)
			tx.Rollback()
			continue
		}

		// Assign user ID to address and save address
		u.Address.UserID = user.ID
		if err := tx.Create(&u.Address).Error; err != nil {
			log.Printf("Failed to seed address for user %s: %v", u.Email, err)
			tx.Rollback()
			continue
		}

		// Commit transaction
		tx.Commit()
		log.Printf("Seeded user and address: %s", u.Email)
	}
}
