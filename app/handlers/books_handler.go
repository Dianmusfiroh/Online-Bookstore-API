package handlers

import (
	"book-online-api/app/dto"
	"book-online-api/app/services"
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookHandler struct {
    service  services.BookService
    validate *validator.Validate
}

func NewBookHandler(service services.BookService) *BookHandler {
    return &BookHandler{
        service:  service,
        validate: validator.New(),
    }
}

// GET /books
func (h *BookHandler) GetAll(c *fiber.Ctx) error {
    var query dto.BookQuery
    if err := c.QueryParser(&query); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
    }

    books, total, err := h.service.FindAll(query)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve books"})
    }

    return c.JSON(fiber.Map{
        "data":  books,
        "total": total,
        "page":  query.Page,
        "limit": query.Limit,
    })
}

// GET /books/:id
func (h *BookHandler) GetByID(c *fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
    }

    book, err := h.service.FindByID(uint(id))
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve book"})
    }
    return c.JSON(book)
}

// POST /books (Admin only)
func (h *BookHandler) Create(c *fiber.Ctx) error {
    var input dto.CreateBookInput
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if err := h.validate.Struct(input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    

    book, err := h.service.Create(input)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "message": book})
    }
    return c.Status(fiber.StatusCreated).JSON(book)
}

// PUT /books/:id (Admin only)
func (h *BookHandler) Update(c *fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
    }
    
    var input dto.UpdateBookInput
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if err := h.validate.Struct(input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    book, err := h.service.Update(uint(id), input)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(book)
}

// DELETE /books/:id (Admin only)
func (h *BookHandler) Delete(c *fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
    }

    if err := h.service.Delete(uint(id)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})

}