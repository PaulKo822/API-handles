package userservice

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
