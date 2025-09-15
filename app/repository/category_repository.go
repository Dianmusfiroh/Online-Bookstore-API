package repository

import (
	"book-online-api/app/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
    Create(category *models.Category) error
    FindAll() ([]models.Category, error)
    FindByID(id uint) (*models.Category, error)
    Update(category *models.Category) error
    Delete(category *models.Category) error
}

type categoryRepository struct {
    db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
    return &categoryRepository{db}
}

func (r *categoryRepository) Create(category *models.Category) error {
    return r.db.Create(category).Error
}

func (r *categoryRepository) FindAll() ([]models.Category, error) {
    var categories []models.Category
    if err := r.db.Find(&categories).Error; err != nil {
        return nil, err
    }
    return categories, nil
}

func (r *categoryRepository) FindByID(id uint) (*models.Category, error) {
    var category models.Category
    if err := r.db.First(&category, id).Error; err != nil {
        return nil, err
    }
    return &category, nil
}

func (r *categoryRepository) Update(category *models.Category) error {
    return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(category *models.Category) error {
    return r.db.Delete(category).Error
}