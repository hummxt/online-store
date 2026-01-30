package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;uniqueIndex" json:"user_id"`
	Items     []CartItem `gorm:"foreignKey:CartID" json:"items"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CartID    uuid.UUID `gorm:"type:uuid" json:"cart_id"`
	ProductID uuid.UUID `gorm:"type:uuid" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity  int       `gorm:"type:int;not null" json:"quantity"`
}

func (c *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

func (ci *CartItem) BeforeCreate(tx *gorm.DB) (err error) {
	ci.ID = uuid.New()
	return
}
