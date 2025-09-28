package main

import (
	"log"
	"os"
	_ "github.com/lib/pq" // Driver untuk "postgresql"
    _ "github.com/jackc/pgx/v5/stdlib" // Driver untuk "pgx"

	"book-online-api/app/routes"
	"book-online-api/config"
	"book-online-api/app/seeder"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

)

func main() {
	// 1. Muat variabel lingkungan lokal
	// Ini penting agar aplikasi bisa berjalan di komputer lokal Anda
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, running with system environment variables.")
	}

	// 2. Inisialisasi Fiber app
	app := fiber.New()
	
	// 3. Hubungkan ke database
	db := config.ConnectDB()
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}

	// 4. Jalankan seeder jika variabel lingkungan disetel
	if os.Getenv("RUN_SEEDER") == "true" {
		seeder.Seed(db)
		log.Println("Seeder executed successfully.")
		os.Exit(0) // Keluar dengan status sukses
	}

	// 5. Tambahkan middleware untuk database dan CORS
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

	// 6. Atur semua rute aplikasi
	routes.SetupRoutes(app, db)
	
	// 7. Dapatkan port dari variabel lingkungan dan jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Port default
	}

	// log.Fatal akan menampilkan pesan error jika server gagal start
	log.Fatal(app.Listen(":" + port))
}