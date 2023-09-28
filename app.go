package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"fmt"
	"api/functions"
	"api/database"
	"api/controllers"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	rep.SetLog()
	log.Info("Création Fiber App")

	// Création de l'application Fiber

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			log.Error("Error code: ", code)
			return ctx.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})
	
	database.InitDatabase()

	// Middleware CORS

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "*",
	}))

	// Controllers

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	
	setupRoutes(app)

	app.Listen(":3000")
}

func setupRoutes(app *fiber.App){

	fmt.Println("GET Products")

    app.Route("/products", func(api fiber.Router){
        api.Get("/",controllers.GetProducts)
        api.Post("/",controllers.AddProduct)
        api.Delete("/",controllers.DeleteProduct)
    })
}
