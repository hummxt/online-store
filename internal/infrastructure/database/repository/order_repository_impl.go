package repository

import (
	"context"
	"ecommerce/internal/domain/entity"
	domrepo "ecommerce/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) domrepo.OrderRepository {
	return &orderRepositoryImpl{db: db}
}

func (r *orderRepositoryImpl) Create(ctx context.Context, order *entity.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Order, error) {
	var order entity.Order
	if err := r.db.WithContext(ctx).Preload("OrderItems.Product").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepositoryImpl) FindByUserID(ctx context.Context, userID uuid.UUID) ([]entity.Order, error) {
	var orders []entity.Order
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepositoryImpl) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	return r.db.WithContext(ctx).Model(&entity.Order{}).Where("id = ?", id).Update("status", status).Error
}
