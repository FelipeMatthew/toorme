package models

import "time"

type TravelPlan struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	CustomerID  uint      `gorm:"not null" json:"customer_id"`
	Customer    User      `gorm:"foreignKey:CustomerID;constraint:onDelete:CASCADE" json:"customer"`
	Status      string    `gorm:"default:'active'" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
