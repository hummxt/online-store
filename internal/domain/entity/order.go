package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID            uuid.UUID   `gorm:"type:uuid;primary_key" json:"id"`
	UserID        uuid.UUID   `gorm:"type:uuid" json:"user_id"`
	TotalAmount   float64     `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	Status        string      `gorm:"type:varchar(20);default:'pending'" json:"status"`
	OrderItems    []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
	PaymentStatus string      `gorm:"type:varchar(20);default:'pending'" json:"payment_status"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	OrderID   uuid.UUID `gorm:"type:uuid" json:"order_id"`
	ProductID uuid.UUID `gorm:"type:uuid" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int       `gorm:"type:int;not null" json:"quantity"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.New()
	return
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	oi.ID = uuid.New()
	return
}
