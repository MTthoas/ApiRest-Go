package controllers

import (
	"github.com/gofiber/fiber/v2"
	"api/database"
)

// GetProducts retrieves all products from the database
func GetProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []database.Product
	db.Find(&products)
	return c.JSON(products)
}

// AddProduct adds a new product to the database
func AddProduct(c *fiber.Ctx) error {
	db := database.DB
	product := new(database.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Failed to parse body",
		})
	}

	db.Create(&product)
	return c.Status(201).JSON(product)
}

// DeleteProduct deletes a product from the database based on its ID
func DeleteProduct(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")

	var product database.Product
	db.First(&product, id)

	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	db.Delete(&product)
	return c.Status(200).JSON(fiber.Map{
		"message": "Product successfully deleted",
	})
}
