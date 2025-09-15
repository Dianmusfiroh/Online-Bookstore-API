package dto

import "book-online-api/app/models"

type OrderItemInput struct {
    BookID   uint `json:"book_id" validate:"required"`
    Quantity int  `json:"quantity" validate:"required,gt=0"`
}

type CreateOrderInput struct {
    Items []OrderItemInput `json:"items" validate:"required,dive"`
}
type CreateOrderResponse struct {
    Order   *models.Order `json:"order"`
    Message string           `json:"message"`
}
