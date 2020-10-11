package models

import (
	"time"
)

type User struct {
	ID int `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
