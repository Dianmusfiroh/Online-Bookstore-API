package handlers

import (
    "book-online-api/app/services"
    "github.com/gofiber/fiber/v2"
)

type ReportHandler struct {
    service services.ReportService
}

func NewReportHandler(service services.ReportService) *ReportHandler {
    return &ReportHandler{service}
}

func (h *ReportHandler) GetSalesReport(c *fiber.Ctx) error {
    revenue, totalSold, err := h.service.GetSalesReport()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get sales report"})
    }
    return c.JSON(fiber.Map{
        "total_revenue":   revenue,
        "total_books_sold": totalSold,
    })
}

func (h *ReportHandler) GetBestsellers(c *fiber.Ctx) error {
    bestsellers, err := h.service.GetBestsellers()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get bestsellers report"})
    }
    return c.JSON(bestsellers)
}

func (h *ReportHandler) GetPriceReport(c *fiber.Ctx) error {
    prices, err := h.service.GetPriceReport()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get price report"})
    }
    return c.JSON(prices)
}