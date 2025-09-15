package handlers

import (
	"book-online-api/app/dto"
	"book-online-api/app/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"errors"
)

type CategoryHandler struct {
    service  services.CategoryService
    validate *validator.Validate
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
    return &CategoryHandler{
        service:  service,
        validate: validator.New(),
    }
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
    var input dto.CreateCategoryInput
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if err := h.validate.Struct(input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    category, err := h.service.Create(input)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create category"})
    }
    return c.Status(fiber.StatusCreated).JSON(category)
}

func (h *CategoryHandler) GetAll(c *fiber.Ctx) error {
    categories, err := h.service.FindAll()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve categories"})
    }
    return c.JSON(categories)
}

func (h *CategoryHandler) GetByID(c *fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
    }

    category, err := h.service.FindByID(uint(id))
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve category"})
    }
    return c.JSON(category)
}

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
    }

    var input dto.UpdateCategoryInput
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if err := h.validate.Struct(input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    category, err := h.service.Update(uint(id), input)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(category)
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
    id, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
    }

    if err := h.service.Delete(uint(id)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Category deleted successfully"})
}