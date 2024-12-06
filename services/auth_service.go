package services

import (
	"errors"
	"task-golang/repositories"
	"task-golang/utils"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func (s *AuthService) Login(email, password string) (string, error) {
	// Find user by email
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Check password
	if !utils.CheckPassword(user.Password, password) {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
