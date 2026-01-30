package repository

import (
	"context"
	"ecommerce/internal/domain/entity"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(ctx context.Context, order *entity.Order) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Order, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Order, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status string) error
}
