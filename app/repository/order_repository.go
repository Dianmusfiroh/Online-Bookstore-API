package repository

import (
	"book-online-api/app/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepository interface {
    Create(order *models.Order) error
    FindByUserID(userID uint) ([]models.Order, error)
    FindByID(orderID uint) (*models.Order, error)
    FindAll() ([]models.Order, error)
    Update(order *models.Order) error
}

type orderRepository struct {
    db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
    return &orderRepository{db}
}

// CreateOrderInTransaction handles the creation of a new order with multiple items
// within a single database transaction. It also handles stock validation and deduction.
func (r *orderRepository) Create(order *models.Order) error {
	// Mulai transaksi database
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	// Pastikan transaksi di-rollback jika terjadi panic
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Kurangi stok buku dan buat order items
	for i, item := range order.OrderItems {
		var book models.Book
		
		// Gunakan .First() di dalam transaksi untuk mengunci baris
		if err := tx.First(&book, item.BookID).Error; err != nil {
			tx.Rollback()
			return errors.New("book not found during transaction")
		}

		// Validasi stok di lapisan service, di sini kita hanya mengurangi
		book.Stock -= item.Quantity
		if err := tx.Save(&book).Error; err != nil {
			tx.Rollback()
			return err
		}
        
        // Perbarui harga di OrderItem dari database
        order.OrderItems[i].Price = book.Price
	}
	fmt.Println("ini order repo",order )
	// Buat pesanan utama
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// Muat ulang (preload) order dari database untuk mendapatkan data relasi lengkap
	if err := tx.Preload("OrderItems.Book").Preload("User").First(order, order.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Preload("OrderItems.Book.Category").Preload("OrderItems.Book").Preload("User").First(order, order.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	

	// Commit transaksi
	return tx.Commit().Error
}

func (r *orderRepository) FindByUserID(userID uint) ([]models.Order, error) {
    var orders []models.Order
    if err := r.db.Preload("OrderItems.Book").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}

func (r *orderRepository) FindByID(orderID uint) (*models.Order, error) {
    var order models.Order
    if err := r.db.Preload("OrderItems.Book").First(&order, orderID).Error; err != nil {
        return nil, err
    }
    return &order, nil
}

func (r *orderRepository) FindAll() ([]models.Order, error) {
    var orders []models.Order
    if err := r.db.Preload("OrderItems.Book").Preload("User").Find(&orders).Error; err != nil {
        return nil, err
    }
    return orders, nil
}

func (r *orderRepository) Update(order *models.Order) error {
	tx := r.db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

		if err := tx.Preload("OrderItems.Book").Preload("User").First(order, order.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Preload("OrderItems.Book.Category").Preload("OrderItems.Book").Preload("User").First(order, order.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Save(order).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
