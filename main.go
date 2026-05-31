// package main

// import (
// 	"log"

// 	db "github.com/gesitnur/go_fiber_step_1/db"

// 	"github.com/gesitnur/go_fiber_step_1/routes"
// 	"github.com/gofiber/fiber/v3"
// )

// // func main() {
// // 	// Di v3, inisialisasi standar menggunakan fiber.New() tetap sama
// // 	app := fiber.New()

// // 	// Endpoint GET /hello
// // 	app.Get("/hello", func(c fiber.Ctx) error { // Di v3, cukup 'fiber.Ctx' tanpa pointer '*'
// // 		return c.SendString("Hello, World!")
// // 	})

// // 	app.Get("/check", func(c fiber.Ctx) error {
// // 		return c.Status(fiber.StatusOK).JSON(fiber.Map{
// // 			"status":  "success",
// // 			"message": "Server is running",
// // 		})
// // 	})
// // 	// Jalankan server di port 3000
// // 	log.Fatal(app.Listen(":3010"))
// // }

// func main() {
// 	err := db.InitDB() // database connection
// 	if err != nil {
// 		log.Fatal("error while connecting to database: ", err)
// 	}
// 	app := fiber.New()
// 	app.Get("/check", func(c fiber.Ctx) error {
// 		m := c.Queries()
// 		message := m["message"]
// 		if message == "" {
// 			message = "default"
// 		}
// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"status":  "200",
// 			"message": message,
// 		})
// 	})
// 	routes.SetupRoutes(app)
// 	err = app.Listen(":8000")
// 	if err != nil {
// 		log.Println("error while starting the server: ", err)
// 	}
// }

package main

import (
	"log"

	"github.com/gesitnur/go_fiber_step_1/controllers"
	"github.com/gesitnur/go_fiber_step_1/initializers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
	app := fiber.New()

	// Global Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		// Change this line to a slice of strings:
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowCredentials: true,
	}))

	// Create an /api route group
	api := app.Group("/api")

	// Attach your routes to the api group
	api.Route("/notes", func(router fiber.Router) {
		// router.Post("/", controllers.CreateNoteHandler)
		router.Get("", controllers.FindNotes)
	})

	// api.Route("/notes/:noteId", func(router fiber.Router) {
	// 	router.Delete("", controllers.DeleteNote)
	// 	router.Get("", controllers.FindNoteById)
	// 	router.Patch("", controllers.UpdateNote)
	// })

	// api.Get("/healthchecker", func(c fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{
	// 		"status":  "success",
	// 		"message": "Welcome to Golang, Fiber, and GORM",
	// 	})
	// })

	log.Fatal(app.Listen(":8000"))
}
