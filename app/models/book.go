package models

import "time"

type Book struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    Title        string    `gorm:"type:varchar(255)" json:"title"`
    Author       string    `gorm:"type:varchar(100)" json:"author"`
    Price        float64   `gorm:"type:decimal(10,2)" json:"price"`
    Stock        int       `gorm:"type:int" json:"stock"`
    Year         int       `gorm:"type:int" json:"year"`
    CategoryID   uint      `json:"category_id"`
    ImageBase64  string    `gorm:"type:text" json:"image_base64"`
    Category     Category  `gorm:"foreignKey:CategoryID"`
    CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}