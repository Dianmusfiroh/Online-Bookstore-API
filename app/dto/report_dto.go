package dto

// SalesReportResponse digunakan untuk endpoint GET /reports/sales
type SalesReportResponse struct {
    TotalRevenue    float64 `json:"total_revenue"`
    TotalBooksSold  int64   `json:"total_books_sold"`
}

// BestsellerBookResponse digunakan untuk endpoint GET /reports/bestseller
type BestsellerBookResponse struct {
    ID        uint   `json:"id"`
    Title     string `json:"title"`
    TotalSold int64  `json:"total_sold"`
}

// PriceReportResponse digunakan untuk endpoint GET /reports/prices
type PriceReportResponse struct {
    MaxPrice float64 `json:"max_price"`
    MinPrice float64 `json:"min_price"`
    AvgPrice float64 `json:"avg_price"`
}