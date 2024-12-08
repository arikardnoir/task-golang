package provider

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"task-golang/models"
)

type ViaCEPProvider struct{}

func (p *ViaCEPProvider) FetchAddress(cep string) (*models.AddressData, error) {
	// Clean the CEP to remove special chars
	cleanedCEP := cleanCEP(cep)
	if len(cleanedCEP) != 8 {
		return nil, errors.New("invalid CEP format")
	}

	// make a request to a API ViaCEP
	apiURL := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cleanedCEP)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, errors.New("failed to fetch address data")
	}
	defer resp.Body.Close()

	// Decode APi response
	var address models.AddressData
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return nil, errors.New("failed to parse address data")
	}

	// Validate is CEP was found
	if address.Street == "" && address.City == "" {
		return nil, errors.New("CEP not found")
	}

	return &address, nil
}

func cleanCEP(cep string) string {
	result := ""
	for _, char := range cep {
		if char >= '0' && char <= '9' {
			result += string(char)
		}
	}
	return result
}
