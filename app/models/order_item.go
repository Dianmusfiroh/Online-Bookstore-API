package models


type OrderItem struct {
    ID       uint    `gorm:"primaryKey" json:"id"`
    OrderID  uint    `json:"order_id"`
    BookID   uint    `json:"book_id"`
    Quantity int     `gorm:"type:int" json:"quantity"`
    Price    float64 `gorm:"type:decimal(10,2)" json:"price"`
    Book     Book    `gorm:"foreignKey:BookID"`

}