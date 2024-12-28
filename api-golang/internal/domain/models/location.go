package models

import "time"

type Location struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	SupplierID  uint      `gorm:"not null" json:"supplier_id"`
	Supplier    Supplier  `gorm:"foreignKey:SupplierID;constraint:onDelete:CASCADE" json:"supplier"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
