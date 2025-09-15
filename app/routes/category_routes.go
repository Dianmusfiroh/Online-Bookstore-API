package routes

import (
	"book-online-api/app/handlers"
	"book-online-api/app/middleware"
	"book-online-api/app/repository"
	"book-online-api/app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterCategoryRoutes(apiRouter *fiber.App, db *gorm.DB) {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	categoryGroup :=apiRouter.Group("/api/categories", middleware.JWTAuth(), middleware.AdminAuth())
	categoryGroup.Get("/", categoryHandler.GetAll)
	categoryGroup.Get("/:id", categoryHandler.GetByID)
	categoryGroup.Put("/:id", categoryHandler.Update)
	categoryGroup.Post("/", categoryHandler.Create)
	categoryGroup.Delete("/:id", categoryHandler.Delete)
}