package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" gorm:"unique" validate:"required,email"`
	Phone string `json:"phone" validate:"required"`
}
