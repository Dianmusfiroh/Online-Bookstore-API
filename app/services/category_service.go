package services

import (
	"book-online-api/app/dto"
	"book-online-api/app/models"
	"book-online-api/app/repository"
	"errors"
)

type CategoryService interface {
    Create(input dto.CreateCategoryInput) (*models.Category, error)
    FindAll() ([]models.Category, error)
    FindByID(id uint) (*models.Category, error)
    Update(id uint, input dto.UpdateCategoryInput) (*models.Category, error)
    Delete(id uint) error
}

type categoryService struct {
    repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
    return &categoryService{repo}
}

func (s *categoryService) Create(input dto.CreateCategoryInput) (*models.Category, error) {
    category := &models.Category{Name: input.Name}
    if err := s.repo.Create(category); err != nil {
        return nil, err
    }
    return category, nil
}

func (s *categoryService) FindAll() ([]models.Category, error) {
    return s.repo.FindAll()
}

func (s *categoryService) FindByID(id uint) (*models.Category, error) {
    return s.repo.FindByID(id)
}

func (s *categoryService) Update(id uint, input dto.UpdateCategoryInput) (*models.Category, error) {
    category, err := s.repo.FindByID(id)
    if err != nil {
        return nil, errors.New("category not found")
    }

    category.Name = input.Name
    if err := s.repo.Update(category); err != nil {
        return nil, err
    }
    return category, nil
}

func (s *categoryService) Delete(id uint) error {
    category, err := s.repo.FindByID(id)
    if err != nil {
        return errors.New("category not found")
    }
    return s.repo.Delete(category)
}