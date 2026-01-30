package repository

import (
	"context"
	"ecommerce/internal/domain/entity"
	domrepo "ecommerce/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type cartRepositoryImpl struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) domrepo.CartRepository {
	return &cartRepositoryImpl{db: db}
}

func (r *cartRepositoryImpl) GetByUserID(ctx context.Context, userID uuid.UUID) (*entity.Cart, error) {
	var cart entity.Cart
	err := r.db.WithContext(ctx).Preload("Items.Product").Where("user_id = ?", userID).FirstOrCreate(&cart, entity.Cart{UserID: userID}).Error
	return &cart, err
}

func (r *cartRepositoryImpl) AddItem(ctx context.Context, item *entity.CartItem) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *cartRepositoryImpl) UpdateItem(ctx context.Context, item *entity.CartItem) error {
	return r.db.WithContext(ctx).Save(item).Error
}

func (r *cartRepositoryImpl) RemoveItem(ctx context.Context, cartID, productID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&entity.CartItem{}).Error
}

func (r *cartRepositoryImpl) ClearCart(ctx context.Context, cartID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("cart_id = ?", cartID).Delete(&entity.CartItem{}).Error
}
