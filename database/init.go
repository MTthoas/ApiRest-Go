package database

import (
	"github.com/gofiber/fiber/v2/log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func InitDatabase(){

	log.Info("Récupération DataBase URL")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL is not set")
	}

	log.Info("Starting server...")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true, 
	}), &gorm.Config{})

	if err != nil {
		log.Error("Error connecting to database: ", err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
	
	db.Create(&User{Name: "John", Age: 20})

	DB = db

	log.Info("Database connected")
}
