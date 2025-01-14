package models

import "time"

type Notification struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"not null" json:"user_id"`
	User      User       `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE" json:"user"`
	Message   string     `gorm:"not null" json:"message"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	ReadAt    *time.Time `json:"read_at"`
}
