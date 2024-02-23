package models

import (
	"time"
)

type Company struct {
	ID        string `gorm:"primary_key" json:"id"`
	Name      string `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Companies []Company
