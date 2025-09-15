package routes

import (
	"book-online-api/app/handlers"
	"book-online-api/app/middleware"
	"book-online-api/app/repository"
	"book-online-api/app/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	// Membuat grup API utama
	api := app.Group("/api")

	// Inisialisasi dependensi
	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Rute publik (tanpa token)
	api.Post("/register", authHandler.Register)
	api.Post("/login", authHandler.Login)

	// Rute yang dilindungi (membutuhkan token JWT)
	protected := api.Group("/", middleware.JWTAuth())
	// Daftarkan rute-rute dari sub-router
	RegisterCategoryRoutes(app, db)
	RegisterOrderRoutes(app, db)
    RegisterBookRoutes(protected, app, db)
	RegisterReportRoutes( app, db)
	
}