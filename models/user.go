package models

import (
	"time"
)

type User struct {
	ID        string `gorm:"primary_key" json:"id"`
	Name      string `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
	Address   string `gorm:"type:text" json:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []User



