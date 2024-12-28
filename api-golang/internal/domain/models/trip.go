package models

import "time"

type Trip struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	CustomerID    uint       `gorm:"not null" json:"customer_id"`
	Customer      User       `gorm:"foreignKey:CustomerID;constraint:onDelete:CASCADE" json:"customer"`
	DriverID      uint       `gorm:"not null" json:"driver_id"`
	Driver        User       `gorm:"foreignKey:DriverID;constraint:onDelete:CASCADE" json:"driver"`
	VehicleID     uint       `json:"vehicle_id"`
	Vehicle       Vehicle    `gorm:"foreignKey:VehicleID;constraint:onDelete:SET NULL" json:"vehicle"`
	TravelPlanID  uint       `json:"travel_plan_id"`
	TravelPlan    TravelPlan `gorm:"foreignKey:TravelPlanID;constraint:onDelete:SET NULL" json:"travel_plan"`
	LocationID    uint       `json:"location_id"`
	Location      Location   `gorm:"foreignKey:LocationID;constraint:onDelete:SET NULL" json:"location"`
	DepartureDate time.Time  `gorm:"not null" json:"departure_date"`
	ReturnDate    *time.Time `json:"return_date"`
	Status        string     `gorm:"default:'pending'" json:"status"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at"`
}
