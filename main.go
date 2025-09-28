package main

import (
	"book-online-api/app/routes"
	"book-online-api/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"book-online-api/app/seeder"
	"os"

)

func main() {
	app := fiber.New()
	db := config.ConnectDB()
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}

	if os.Getenv("RUN_SEEDER") == "true" {
		seeder.Seed(db)
		log.Println("Seeder executed, exiting.")
	}
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers",
		AllowCredentials: false,
	}))

	routes.SetupRoutes(app, db)
	// app.Listen(":3000")
	log.Fatal("Failed to start the server: ", app.Listen(":3000"))
}