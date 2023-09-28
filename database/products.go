package database

import (
	// "gorm.io/gorm"
)

type Product struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"type:varchar(100)"`
    Price float64
}