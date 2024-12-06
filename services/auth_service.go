package services

import (
	"errors"
	"math/rand"
	"task-golang/repositories"
	"task-golang/utils"
	"time"
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

func (s *AuthService) GenerateRecoveryToken(email string) (string, error) {
	// Find user by email
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("email not found")
	}

	// Generate recovery token
	token := generateToken()
	expiration := time.Now().Add(1 * time.Hour) // Token valid for 1 hour

	// Save token to user
	err = s.UserRepo.SetRecoveryToken(user, token, expiration)
	if err != nil {
		return "", errors.New("could not set recovery token")
	}

	// Mock sending email (for now, just log the token)
	MockSendEmail(user.Email, token)

	return token, nil
}

func generateToken() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, 20)
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}
	return string(token)
}

func MockSendEmail(email, token string) {
	// Replace with actual email service
	println("Mock email sent to:", email)
	println("Recovery token:", token)
}
