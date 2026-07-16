package models

import (
	"time"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Price       int       `json:"price" gorm:"not null"`
	Stock       int       `json:"stock" gorm:"not null"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
