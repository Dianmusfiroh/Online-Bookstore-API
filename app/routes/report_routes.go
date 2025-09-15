package routes

import (
	"book-online-api/app/handlers"
	"book-online-api/app/middleware"
	"book-online-api/app/repository"
	"book-online-api/app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterReportRoutes(apiRouter *fiber.App, db *gorm.DB) {
	reportRepo := repository.NewReportRepository(db)
	reportService := services.NewReportService(reportRepo)
	reportHandler := handlers.NewReportHandler(reportService)
	reportGroup :=apiRouter.Group("/api/reports", middleware.JWTAuth(), middleware.AdminAuth())
	reportGroup.Get("/sales", reportHandler.GetSalesReport)
    reportGroup.Get("/bestseller", reportHandler.GetBestsellers)
    reportGroup.Get("/prices", reportHandler.GetPriceReport)
}