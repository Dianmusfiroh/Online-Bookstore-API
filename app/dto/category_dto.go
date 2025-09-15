package dto

type CreateCategoryInput struct {
    Name string `json:"name" validate:"required,min=3"`
}

type UpdateCategoryInput struct {
    Name string `json:"name" validate:"required,min=3"`
}