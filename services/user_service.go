package services

import (
	"errors"
	"task-golang/models"
	"task-golang/repositories"
	"task-golang/utils"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.UserRepo.GetAllUsers()
}

func (s *UserService) Register(user *models.User) error {
	// Check if email is already registered
	_, err := s.UserRepo.FindByEmail(user.Email)
	if err == nil {
		return errors.New("email already in use")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Save user
	return s.UserRepo.Create(user)
}
