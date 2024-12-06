package repositories

import (
	"github.com/google/uuid"
	"task-golang/database"
	"task-golang/models"
)

type UserRepository struct{}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *UserRepository) FindByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
