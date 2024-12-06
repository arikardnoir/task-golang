package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	Address   string    `gorm:"type:text"`
	CreatedAt int64     `gorm:"autoCreateTime"`
	UpdatedAt int64     `gorm:"autoUpdateTime"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
