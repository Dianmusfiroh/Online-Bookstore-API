package routes

import (
    "book-online-api/app/handlers"
    "book-online-api/app/middleware" // Impor middleware di sini
    "book-online-api/app/repository"
    "book-online-api/app/services"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

// dan router utama untuk membuat grup admin.
func RegisterBookRoutes(protectedRouter fiber.Router, apiRouter *fiber.App, db *gorm.DB) {

    bookRepo := repository.NewBookRepository(db)
    bookService := services.NewBookService(bookRepo)
    bookHandler := handlers.NewBookHandler(bookService)

    // Rute yang dapat diakses oleh user dan admin (GET)
    protectedRouter.Get("/books", bookHandler.GetAll)
    protectedRouter.Get("/books/:id", bookHandler.GetByID)

    // Grup rute KHUSUS ADMIN untuk buku
    // Middleware AdminAuth diterapkan langsung di sini
    adminBookRouter := apiRouter.Group("/api/books", middleware.JWTAuth(), middleware.AdminAuth())
    
    // Rute admin-only
    adminBookRouter.Post("/", bookHandler.Create)
    adminBookRouter.Put("/:id", bookHandler.Update)
    adminBookRouter.Delete("/:id", bookHandler.Delete)
}