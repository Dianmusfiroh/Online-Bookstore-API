package repository

import (
    "book-online-api/app/models"
    "gorm.io/gorm"
)

type ReportRepository interface {
    GetSalesReport() (float64, int64, error)
    GetBestsellers() ([]models.BestsellerBook, error)
    GetPriceReport() (*models.PriceReport, error)
}

type reportRepository struct {
    db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
    return &reportRepository{db}
}

// GetSalesReport mengambil total omzet (dari order yang PAID) dan total buku terjual
func (r *reportRepository) GetSalesReport() (float64, int64, error) {
    var totalRevenue float64
    var totalBooksSold int64

    // Hitung total omzet dari pesanan yang sudah dibayar
    r.db.Model(&models.Order{}).Where("status = ?", "PAID").Select("sum(total_price)").Scan(&totalRevenue)

    // Hitung total buku terjual dari item pesanan yang sudah dibayar
    r.db.Model(&models.OrderItem{}).
        Joins("JOIN orders ON orders.id = order_items.order_id").
        Where("orders.status = ?", "PAID").
        Select("sum(quantity)").Scan(&totalBooksSold)

    return totalRevenue, totalBooksSold, nil
}

// GetBestsellers mengambil 3 buku terlaris
func (r *reportRepository) GetBestsellers() ([]models.BestsellerBook, error) {
    var bestsellers []models.BestsellerBook
    
    // Query untuk mengambil 3 buku terlaris berdasarkan total kuantitas terjual
    if err := r.db.Model(&models.OrderItem{}).
        Select("books.id, books.title, sum(order_items.quantity) as total_sold").
        Joins("JOIN orders ON orders.id = order_items.order_id").
        Joins("JOIN books ON books.id = order_items.book_id").
        Where("orders.status = ?", "PAID").
        Group("books.id, books.title").
        Order("total_sold desc").
        Limit(3).
        Scan(&bestsellers).Error; err != nil {
        return nil, err
    }
    
    return bestsellers, nil
}

// GetPriceReport mengambil harga maksimum, minimum, dan rata-rata buku
func (r *reportRepository) GetPriceReport() (*models.PriceReport, error) {
    var priceReport models.PriceReport

    // Gunakan GORM untuk menghitung nilai agregat
    if err := r.db.Model(&models.Book{}).
        Select("max(price) as max_price, min(price) as min_price, avg(price) as avg_price").
        Scan(&priceReport).Error; err != nil {
        return nil, err
    }

    return &priceReport, nil
}