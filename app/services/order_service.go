package services

import (
	"book-online-api/app/dto"
	"book-online-api/app/models"
	"book-online-api/app/repository"
	"context"
	"encoding/json"
	"os"
	"time"

	"errors"
	"fmt"


	"github.com/shopspring/decimal"
	xendit "github.com/xendit/xendit-go/v7"
	invoice "github.com/xendit/xendit-go/v7/invoice"
)

type OrderService interface {
    // baru
    // Create(userID uint, input dto.CreateOrderInput) (*models.Order, string, error)
    // lama
    Create(userID uint, input dto.CreateOrderInput) (*models.Order, error)
    Pay(orderID uint) (*models.Order, error)
    FindUserOrders(userID uint) ([]models.Order, error)
    FindAllOrders() ([]models.Order, error)
    FindByID(orderID uint) (*models.Order, error)
}

type orderService struct {
    orderRepo repository.OrderRepository
    bookRepo  repository.BookRepository // Tambahkan bookRepo

}


// func NewOrderService(orderRepo repository.OrderRepository, bookRepo repository.BookRepository) (*orderService, error) {

func NewOrderService(orderRepo repository.OrderRepository, bookRepo repository.BookRepository) OrderService {
     return &orderService{
        orderRepo: orderRepo,
        bookRepo:  bookRepo,
    }   
}

func (s *orderService) Create(userID uint, input dto.CreateOrderInput  ) (*models.Order, error) {
    var totalPrice float64
    var bookName string
    var allXndt string
    orderItems := make([]models.OrderItem, len(input.Items))

    // Validasi input awal dan hitung total harga
    for i, itemInput := range input.Items {
        // Cek apakah buku ada di database
        book, err := s.bookRepo.FindByID(itemInput.BookID)
        if err != nil {
            return nil, errors.New("book not found")
        }
        
        // Validasi stok
        if book.Stock < itemInput.Quantity {
            return nil, fmt.Errorf("stock of book with id %d is not enough", itemInput.BookID)
        }

        orderItems[i] = models.OrderItem{
            BookID:   book.ID,
            Quantity: itemInput.Quantity,
            Price:    book.Price,
        }
        totalPrice += book.Price * float64(itemInput.Quantity)
        bookName = book.Title
    }
   
    // Panggil repository untuk membuat pesanan di dalam transaksi
   
secretKey := os.Getenv("XENDIT_SECRET_KEY")
    if secretKey == "" {
        return nil, errors.New("XENDIT_SECRET_KEY environment variable is not set")
    }
    amountFloat, _ := decimal.NewFromFloat(totalPrice).Float64()

    xndClient := xendit.NewClient(secretKey)
    createInvoiceRequest := &invoice.CreateInvoiceRequest{
        ExternalId: fmt.Sprintf("payment-%d-%d-%s", time.Now().Unix(),userID,   bookName),
    Amount: amountFloat,
    }
    // Pastikan klien berhasil diinisialisasi
  resp, r, err := xndClient.InvoiceApi.CreateInvoice(context.Background()).
        CreateInvoiceRequest(*createInvoiceRequest).
        Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoiceApi.CreateInvoice``: %v\n", err.Error())

        b, _ := json.Marshal(err.FullError())
        fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
        fmt.Fprintf(os.Stdout, "Response from `InvoiceApi.CreateInvoice`: %v\n", resp)
        allXndt = fmt.Sprintf("%v", resp)

    order := &models.Order{
        UserID:     userID,
        TotalPrice: totalPrice,
        Status:     "PENDING",
        OrderItems: orderItems,
        XenditData: allXndt,
        InvoiceURL: resp.InvoiceUrl,
    }
   
     if err := s.orderRepo.Create(order); err != nil {
        return nil, err
    }
    
    return order, nil
}


func (s *orderService) Pay(orderID uint) (*models.Order, error) {
    order, err := s.orderRepo.FindByID(orderID)
    if err != nil {
        return nil, errors.New("order not found")
    }

    if order.Status != "PENDING" {
        return nil, errors.New("order cannot be paid, current status is " + order.Status)
    }

  

    order.Status = "PAID"
    if err := s.orderRepo.Update(order); err != nil {
        return nil, err
    }
    
    return order, nil
}

func (s *orderService) FindUserOrders(userID uint) ([]models.Order, error) {
    return s.orderRepo.FindByUserID(userID)
}

func (s *orderService) FindAllOrders() ([]models.Order, error) {
    return s.orderRepo.FindAll()
}

func (s *orderService) FindByID(orderID uint) (*models.Order, error) {
    return s.orderRepo.FindByID(orderID)
}