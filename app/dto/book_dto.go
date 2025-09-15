package dto

type CreateBookInput struct {
    Title       string  `json:"title" validate:"required"`
    Author      string  `json:"author" validate:"required"`
    Price       float64 `json:"price" validate:"required,gt=0"`
    Stock       int     `json:"stock" validate:"required,gte=0"`
    Year        int     `json:"year" validate:"required,gt=1000,lte=9999"`
    CategoryID  uint    `json:"category_id" validate:"required"`
    ImageBase64 string  `json:"image_base64"`
}

type UpdateBookInput struct {
    Title       string  `json:"title" validate:"required"`
    Author      string  `json:"author" validate:"required"`
    Price       float64 `json:"price" validate:"required,gt=0"`
    Stock       int     `json:"stock" validate:"required,gte=0"`
    Year        int     `json:"year" validate:"required,gt=1000,lte=9999"`
    CategoryID  uint    `json:"category_id" validate:"required"`
    ImageBase64 string  `json:"image_base64"`
}

type BookQuery struct {
    Page       int    `query:"page"`
    Limit      int    `query:"limit"`
    Search     string `query:"search"`
    CategoryID uint   `query:"category_id"`
}