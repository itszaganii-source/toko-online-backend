package models

import (
	"time"
)

type Order struct {
	ID           uint        `json:"id" gorm:"primaryKey"`
	CustomerName string      `json:"customer_name" gorm:"not null"`
	TotalPrice   int         `json:"total_price" gorm:"not null"`
	Status       string      `json:"status" gorm:"default:'Pending'"`
	OrderItems   []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	Price     int     `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
