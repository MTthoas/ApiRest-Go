package controllers

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/log"
	"fmt"
	// "os"
	// "api/database"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func GetUsers(c *fiber.Ctx) error {
	fmt.Println("Get /users from ", c.IP())
	return c.SendString("Hello, World!")
}


	