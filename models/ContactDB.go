package models

import (
	"time"
)

type Contact struct {
	ID int `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Subject string `gorm:"column:subject"`
	Message string `gorm:"column:message"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
