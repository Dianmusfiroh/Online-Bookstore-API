package services

import (
	"book-online-api/app/dto"
	"book-online-api/app/models"
	"book-online-api/app/repository"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BookService interface {
    Create(input dto.CreateBookInput) (*models.Book, error)
    FindAll(query dto.BookQuery) ([]models.Book, int64, error)
    FindByID(id uint) (*models.Book, error)
    Update(id uint, input dto.UpdateBookInput) (*models.Book, error)
    Delete(id uint) error
}

type bookService struct {
    repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
    return &bookService{repo}
}

func (s *bookService) Create(input dto.CreateBookInput) (*models.Book, error) {
    book := &models.Book{
        Title:       input.Title,
        Author:      input.Author,
        Price:       input.Price,
        Stock:       input.Stock,
        Year:        input.Year,
        CategoryID:  input.CategoryID,
        ImageBase64: input.ImageBase64,
    }

    err := s.repo.Create(book)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            if _, err := s.repo.FindCategoryByID(input.CategoryID); err != nil {
                return nil, fmt.Errorf("category with id %d not found", input.CategoryID)
            }
        }
        return nil, err
    }
    return book, nil
}

func (s *bookService) FindAll(query dto.BookQuery) ([]models.Book, int64, error) {
	return s.repo.FindAll(&query)
}

func (s *bookService) FindByID(id uint) (*models.Book, error) {
    return s.repo.FindByID(id)
}

func (s *bookService) Update(id uint, input dto.UpdateBookInput) (*models.Book, error) {
    book, err := s.repo.FindByID(id)
    if err != nil {
        return nil, errors.New("book not found")
    }

    book.Title = input.Title
    book.Author = input.Author
    book.Price = input.Price
    book.Stock = input.Stock
    book.Year = input.Year
    book.CategoryID = input.CategoryID
    book.ImageBase64 = input.ImageBase64
    
    if err := s.repo.Update(book); err != nil {
        return nil, err
    }
    return book, nil
}

func (s *bookService) Delete(id uint) error {
    book, err := s.repo.FindByID(id)
    if err != nil {
        return errors.New("book not found")
    }
    return s.repo.Delete(book)
}