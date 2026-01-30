package repository

import (
	"context"
	"ecommerce/internal/domain/entity"

	"github.com/google/uuid"
)

type CartRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.Cart, error)
	AddItem(ctx context.Context, item *entity.CartItem) error
	UpdateItem(ctx context.Context, item *entity.CartItem) error
	RemoveItem(ctx context.Context, cartID, productID uuid.UUID) error
	ClearCart(ctx context.Context, cartID uuid.UUID) error
}
