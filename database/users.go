package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
}
