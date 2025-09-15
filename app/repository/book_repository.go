package repository

import (
    "book-online-api/app/models"
    "book-online-api/app/dto"
    "gorm.io/gorm"
)

type BookRepository interface {
    Create(book *models.Book) error
    FindAll(query *dto.BookQuery) ([]models.Book, int64, error)
    FindByID(id uint) (*models.Book, error)
    Update(book *models.Book) error
    Delete(book *models.Book) error
    FindCategoryByID(id uint) (*models.Category, error)
}

type bookRepository struct {
    db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
    return &bookRepository{db}
}

func (r *bookRepository) Create(book *models.Book) error {
    return r.db.Create(book).Error
}

func (r *bookRepository) FindAll(query *dto.BookQuery) ([]models.Book, int64, error) {
    var books []models.Book
    var count int64

    tx := r.db.Model(&models.Book{})

    if query.Search != "" {
        tx = tx.Where("title ILIKE ? OR author ILIKE ?", "%"+query.Search+"%", "%"+query.Search+"%")
    }
    
    if query.CategoryID > 0 {
        tx = tx.Where("category_id = ?", query.CategoryID)
    }

    if err := tx.Count(&count).Error; err != nil {
        return nil, 0, err
    }

    if query.Page == 0 {
        query.Page = 1
    }
    if query.Limit == 0 {
        query.Limit = 10
    }

    offset := (query.Page - 1) * query.Limit
    if err := tx.Limit(query.Limit).Offset(offset).Find(&books).Error; err != nil {
        return nil, 0, err
    }

     if err := tx.Preload("Category").Limit(query.Limit).Offset(offset).Find(&books).Error; err != nil {
        return nil, 0, err
    }
    
    return books, count, nil
}

func (r *bookRepository) FindByID(id uint) (*models.Book, error) {
    var book models.Book
    if err := r.db.Preload("Category").First(&book, id).Error; err != nil {
        return nil, err
    }
    return &book, nil
}

func (r *bookRepository) Update(book *models.Book) error {
    return r.db.Save(book).Error
}

func (r *bookRepository) Delete(book *models.Book) error {
    return r.db.Delete(book).Error
}
func (r *bookRepository) FindCategoryByID(id uint) (*models.Category, error) {
    var category models.Category
    if err := r.db.Preload("Book.Category").First(&category, id).Error; err != nil {
        return nil, err
    }
    return &category, nil
}
