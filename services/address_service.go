package services

import (
	"task-golang/models"
	"task-golang/provider"
)

type AddressService struct {
	Provider provider.ViaCEPProvider
}

func (a *AddressService) ValidateAndFetchAddress(cep string) (*models.AddressData, error) {
	return a.Provider.FetchAddress(cep)
}
