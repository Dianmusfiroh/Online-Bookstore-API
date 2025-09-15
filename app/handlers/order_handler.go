package handlers

import (
	"book-online-api/app/dto"
	"book-online-api/app/middleware"
	"book-online-api/app/services"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type OrderHandler struct {
    service services.OrderService
    validate *validator.Validate
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
    return &OrderHandler{
        service:  service,
        validate: validator.New(),
    }
}

// POST /orders
func (h *OrderHandler) Create(c *fiber.Ctx) error {
    claims, ok := c.Locals("claims").(*middleware.JwtClaims)
    if !ok {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user ID from token"})
    }
    userID := claims.UserID

    var input dto.CreateOrderInput
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if err := h.validate.Struct(input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    
    order, err := h.service.Create(userID, input )
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(fiber.StatusCreated).JSON(order)
}

// POST /orders/:id/pay
func (h *OrderHandler) Pay(c *fiber.Ctx) error {
    orderID, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
    }
    
    order, err := h.service.Pay(uint(orderID))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Payment successful", "order": order})
}

// GET /orders
func (h *OrderHandler) GetOrders(c *fiber.Ctx) error {
    claims, ok := c.Locals("claims").(*middleware.JwtClaims)
    role := claims.Role

    if !ok {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user role"})
    }
    if role == "admin" {
        orders, err := h.service.FindAllOrders()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve orders"})
        }
        return c.JSON(orders)
    } else {
        userID := claims.UserID

        orders, err := h.service.FindUserOrders(userID)
        fmt.Printf("[Handler] retrieved %d orders for user %d\n", len(orders), err) // Debugging line
       
        if err != nil {
            
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("Failed to retrieve user %d's orders", userID)})
        }
        if len(orders) == 0 {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User has no orders"})
        }
        return c.JSON(orders)
    }
}
    


// GET /orders/:id
func (h *OrderHandler) GetByID(c *fiber.Ctx) error {
    orderID, err := strconv.ParseUint(c.Params("id"), 10, 32)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid order ID"})
    }
    claims, ok := c.Locals("claims").(*middleware.JwtClaims)
    if !ok {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user ID from token"})
    }
   
    userID := claims.UserID
    role := claims.Role

    order, err := h.service.FindByID(uint(orderID))
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve order"})
    }
    
    if role != "admin" && order.UserID != userID {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: Not the owner or admin"})
    }
    
    return c.JSON(order)
}