package main

import (
	auth "fiber/controllers/auth"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendStatus(fiber.StatusOK)

		c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "You've successfully written to the client",
		})

		return nil
	})

	app.Post("/createUser", auth.CreateUser)

	app.Listen(":8001")
}
