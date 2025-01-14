package models

import "time"

type Vehicle struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	DriverID     uint      `gorm:"not null" json:"driver_id"`
	Driver       User      `gorm:"foreignKey:DriverID;constraint:onDelete:CASCADE" json:"driver"`
	VehicleType  string    `json:"vehicle_type"`
	LicensePlate string    `gorm:"unique" json:"license_plate"`
	Capacity     int       `json:"capacity"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}
