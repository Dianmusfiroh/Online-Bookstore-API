package models

// BestsellerBook digunakan untuk menyimpan data buku terlaris
type BestsellerBook struct {
    ID        uint   `json:"id"`
    Title     string `json:"title"`
    TotalSold int64  `json:"total_sold"`
}

// PriceReport digunakan untuk menyimpan data harga maksimum, minimum, dan rata-rata
type PriceReport struct {
    MaxPrice float64 `json:"max_price"`
    MinPrice float64 `json:"min_price"`
    AvgPrice float64 `json:"avg_price"`
}