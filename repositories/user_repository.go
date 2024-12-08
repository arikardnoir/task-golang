package repositories

import (
	"github.com/google/uuid"
	"task-golang/database"
	"task-golang/models"
	"time"
)

type UserRepository struct{}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetAllUsersWithAddress() ([]models.User, error) {
	var users []models.User
	err := database.DB.Preload("Address").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

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

func (r *UserRepository) SetRecoveryToken(user *models.User, token string, expiration time.Time) error {
	user.RecoveryToken = token
	user.RecoveryTokenExp = expiration
	return database.DB.Save(user).Error
}

func (r *UserRepository) FindByRecoveryToken(token string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("recovery_token = ?", token).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdatePassword(user *models.User, newPassword string) error {
	user.Password = newPassword
	user.RecoveryToken = ""             // Clear recovery token
	user.RecoveryTokenExp = time.Time{} // Clear expiration time
	return database.DB.Save(user).Error
}
