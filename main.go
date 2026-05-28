package main

import (
	"log"

	db "github.com/gesitnur/go_fiber_step_1/db"

	"github.com/gesitnur/go_fiber_step_1/routes"
	"github.com/gofiber/fiber/v3"
)

// func main() {
// 	// Di v3, inisialisasi standar menggunakan fiber.New() tetap sama
// 	app := fiber.New()

// 	// Endpoint GET /hello
// 	app.Get("/hello", func(c fiber.Ctx) error { // Di v3, cukup 'fiber.Ctx' tanpa pointer '*'
// 		return c.SendString("Hello, World!")
// 	})

// 	app.Get("/check", func(c fiber.Ctx) error {
// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"status":  "success",
// 			"message": "Server is running",
// 		})
// 	})
// 	// Jalankan server di port 3000
// 	log.Fatal(app.Listen(":3010"))
// }

func main() {
	err := db.InitDB() // database connection
	if err != nil {
		log.Fatal("error while connecting to database: ", err)
	}
	app := fiber.New()
	app.Get("/check", func(c fiber.Ctx) error {
		m := c.Queries()
		message := m["message"]
		if message == "" {
			message = "default"
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "200",
			"message": message,
		})
	})
	routes.SetupRoutes(app)
	err = app.Listen(":8000")
	if err != nil {
		log.Println("error while starting the server: ", err)
	}
}
