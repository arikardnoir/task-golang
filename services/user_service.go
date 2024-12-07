package services

import (
	"errors"
	"task-golang/database"
	"task-golang/models"
	"task-golang/provider"
	"task-golang/repositories"
	"task-golang/utils"
)

type UserService struct {
	UserRepo    *repositories.UserRepository
	AddressRepo *repositories.AddressRepository
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.UserRepo.GetAllUsers()
}

func (s *UserService) Register(req *models.RegisterRequest) error {
	// Validate email uniqueness
	_, err := s.UserRepo.FindByEmail(req.Email)
	if err == nil {
		return errors.New("email already in use")
	}

	// Usar o AddressService para buscar os dados do endereço
	addressService := AddressService{Provider: provider.ViaCEPProvider{}}
	address, err := addressService.ValidateAndFetchAddress(req.CEP)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Save user and address in a transaction
	tx := database.DB.Begin()

	if err := s.UserRepo.Create(user); err != nil {
		tx.Rollback()
		return err
	}

	// Criar o endereço associado ao usuário
	addressEntity := &models.Address{
		Street:       address.Street,
		Neighborhood: address.Neighborhood,
		Number:       address.Number,
		City:         address.City,
		State:        address.State,
		CEP:          req.CEP,
		UserID:       user.ID,
	}

	addressEntity.UserID = user.ID
	if err := s.AddressRepo.Create(addressEntity); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
