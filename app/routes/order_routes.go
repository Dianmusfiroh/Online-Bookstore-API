package routes

import (
	"book-online-api/app/handlers"
	"book-online-api/app/middleware"
	"book-online-api/app/repository"
	"book-online-api/app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterOrderRoutes(apiRouter *fiber.App, db *gorm.DB) {
	corderRepo := repository.NewOrderRepository(db)
	bookRepo := repository.NewBookRepository(db)
	orderService := services.NewOrderService(corderRepo, bookRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	orderGroup :=apiRouter.Group("/api/orders", middleware.JWTAuth())
	orderGroup.Get("/", orderHandler.GetOrders)
    orderGroup.Post("/", orderHandler.Create)
	orderGroup.Get("/:id", orderHandler.GetByID)
    orderGroup.Post("/:id/pay", orderHandler.Pay)
}