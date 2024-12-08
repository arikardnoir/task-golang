package repositories

import (
	"task-golang/database"
	"task-golang/models"
)

type AddressRepository struct{}

func (r *AddressRepository) Create(address *models.Address) error {
	return database.DB.Create(address).Error
}
