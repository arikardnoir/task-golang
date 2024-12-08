package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Street       string    `gorm:"type:varchar(200);not null"`
	Neighborhood string    `gorm:"type:varchar(100);not null"`
	Number       string    `gorm:"type:varchar(20);not null"`
	City         string    `gorm:"type:varchar(100);not null"`
	State        string    `gorm:"type:varchar(50);not null"`
	CEP          string    `gorm:"type:varchar(20);not null"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;index"`
}

type AddressData struct {
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	Number       string `json:"unidade"`
}

func (a *Address) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
