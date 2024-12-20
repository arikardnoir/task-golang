package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name             string    `gorm:"type:varchar(100);not null"`
	Email            string    `gorm:"type:varchar(100);unique;not null"`
	Password         string    `gorm:"type:varchar(100);not null"`
	Address          Address   `gorm:"foreignKey:UserID"`
	RecoveryToken    string    `gorm:"type:varchar(100);"`
	RecoveryTokenExp time.Time `gorm:"type:datetime;"`
	CreatedAt        int64     `gorm:"autoCreateTime"`
	UpdatedAt        int64     `gorm:"autoUpdateTime"`
}

type RegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
	CEP             string `json:"cep" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
