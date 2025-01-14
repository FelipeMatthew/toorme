package models

import "time"

type Supplier struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	ContactEmail string    `json:"contact_email"`
	ContactPhone string    `json:"contact_phone"`
	Address      string    `json:"address"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
