package service

import (
	"context"
	"ecommerce/internal/domain/entity"
	"ecommerce/internal/domain/repository"
	"errors"

	"github.com/google/uuid"
)

type CartService interface {
	GetCart(ctx context.Context, userID uuid.UUID) (*entity.Cart, error)
	AddToCart(ctx context.Context, userID, productID uuid.UUID, quantity int) error
	UpdateCartItem(ctx context.Context, userID, productID uuid.UUID, quantity int) error
	RemoveFromCart(ctx context.Context, userID, productID uuid.UUID) error
}

type cartService struct {
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewCartService(cartRepo repository.CartRepository, productRepo repository.ProductRepository) CartService {
	return &cartService{cartRepo: cartRepo, productRepo: productRepo}
}

func (s *cartService) GetCart(ctx context.Context, userID uuid.UUID) (*entity.Cart, error) {
	return s.cartRepo.GetByUserID(ctx, userID)
}

func (s *cartService) AddToCart(ctx context.Context, userID, productID uuid.UUID, quantity int) error {
	cart, err := s.cartRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	product, err := s.productRepo.FindByID(ctx, productID)
	if err != nil {
		return errors.New("product not found")
	}

	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	for _, item := range cart.Items {
		if item.ProductID == productID {
			item.Quantity += quantity
			return s.cartRepo.UpdateItem(ctx, &item)
		}
	}

	newItem := &entity.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		Quantity:  quantity,
	}
	return s.cartRepo.AddItem(ctx, newItem)
}

func (s *cartService) UpdateCartItem(ctx context.Context, userID, productID uuid.UUID, quantity int) error {
	cart, err := s.cartRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}

	for _, item := range cart.Items {
		if item.ProductID == productID {
			if quantity <= 0 {
				return s.cartRepo.RemoveItem(ctx, cart.ID, productID)
			}
			item.Quantity = quantity
			return s.cartRepo.UpdateItem(ctx, &item)
		}
	}
	return errors.New("item not found in cart")
}

func (s *cartService) RemoveFromCart(ctx context.Context, userID, productID uuid.UUID) error {
	cart, err := s.cartRepo.GetByUserID(ctx, userID)
	if err != nil {
		return err
	}
	return s.cartRepo.RemoveItem(ctx, cart.ID, productID)
}
