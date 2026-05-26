package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Di v3, inisialisasi standar menggunakan fiber.New() tetap sama
	app := fiber.New()

	// Endpoint GET /hello
	app.Get("/hello", func(c fiber.Ctx) error { // Di v3, cukup 'fiber.Ctx' tanpa pointer '*'
		return c.SendString("Hello, World!")
	})

	app.Get("/check", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Server is running",
		})
	})
	// Jalankan server di port 3000
	log.Fatal(app.Listen(":3010"))
}
