package models

import "time"

type Order struct {
    ID         uint      `gorm:"primaryKey" json:"id"`
    UserID     uint      `json:"user_id"`
    User       User      `gorm:"foreignKey:UserID"`
    TotalPrice float64   `gorm:"type:decimal(10,2)" json:"total_price"`
    Status     string    `gorm:"type:enum('PENDING','PAID','CANCELLED');default:'PENDING'" json:"status"`
    CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
    OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
    XenditData string `gorm:"type:text" json:"xendit_data"` // Simpan semua data dari Xendit
    InvoiceURL string    `gorm:"type:text" json:"invoice_url"` // Tambahkan field untuk menyimpan URL invoice
}

