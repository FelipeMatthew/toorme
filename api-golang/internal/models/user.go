package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Username    string    `gorm:"unique;not null" json:"username"`
	Password    string    `gorm:"not null" json:"password"`
	FullName    string    `json:"full_name"`
	Email       string    `gorm:"unique" json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `gorm:"not null;check:role IN ('admin','customer','driver')" json:"role"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
