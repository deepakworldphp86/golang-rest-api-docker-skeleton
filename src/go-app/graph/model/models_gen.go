package model

import (
	"time"
)

type Categories struct {
	ID           uint      `gorm:"primaryKey"`
	CategoryName string    `gorm:"column:category_name"`
	Description  string    `gorm:"column:description"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

type NewCategory struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}
