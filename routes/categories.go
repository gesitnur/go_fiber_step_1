package routes

import (
	"log"

	db "github.com/gesitnur/go_fiber_step_1/db"
	"github.com/gesitnur/go_fiber_step_1/models"

	fiber "github.com/gofiber/fiber/v3"
)

func GetAllCategories(c fiber.Ctx) error {
	if db.DB == nil {
		return c.Status(500).SendString("database connection is not initialized")
	}
	rows, err := db.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		log.Println("error while running the query on the database: ", err)
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var categories []models.Categories
	for rows.Next() {
		var category models.Categories
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Println("error while scanning the rows: ", err)
			return c.Status(500).SendString(err.Error())
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(categories)
}

func AddCategory(c fiber.Ctx) error {
	category := new(models.Categories)
	if err := c.Bind().JSON(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	query := `INSERT INTO categories (name) VALUES ($1) RETURNING id`
	err := db.DB.QueryRow(query, category.Name).Scan(&category.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create category"})
	}

	return c.Status(200).JSON(category)
}

func SetupRoutes(app *fiber.App) {
	app.Get("/categories", GetAllCategories)
	app.Post("/category", AddCategory)
}
