package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Text   string `json:"text"`
	IsDone bool   `json:"is_done,omitempty"`
	UserID uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Tasks    []Task `gorm:"foreignKey:UserID"`
}
