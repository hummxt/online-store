package dto

import "github.com/google/uuid"

type AddToCartRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required,min=1"`
}

type UpdateCartItemRequest struct {
	Quantity int `json:"quantity" binding:"required,min=0"`
}

type PlaceOrderResponse struct {
	OrderID     uuid.UUID `json:"order_id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
}
