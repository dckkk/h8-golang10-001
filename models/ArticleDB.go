package models

import (
	"time"
)

type Article struct {
	ID int `gorm:"AUTO_INCREMENT"`
	Title string `gorm:"column:title"`
	Text string `gorm:"type:text"`
	Publish string `gorm:"column:publish"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
