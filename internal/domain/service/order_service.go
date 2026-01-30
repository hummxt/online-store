package service

import (
	"context"
	"ecommerce/internal/domain/entity"
	"ecommerce/internal/domain/repository"
	"errors"

	"github.com/google/uuid"
)

type OrderService interface {
	PlaceOrder(ctx context.Context, userID uuid.UUID) (*entity.Order, error)
	GetOrder(ctx context.Context, orderID uuid.UUID) (*entity.Order, error)
	GetUserOrders(ctx context.Context, userID uuid.UUID) ([]entity.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error
}

type orderService struct {
	orderRepo   repository.OrderRepository
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewOrderService(orderRepo repository.OrderRepository, cartRepo repository.CartRepository, productRepo repository.ProductRepository) OrderService {
	return &orderService{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *orderService) PlaceOrder(ctx context.Context, userID uuid.UUID) (*entity.Order, error) {
	cart, err := s.cartRepo.GetByUserID(ctx, userID)
	if err != nil || len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	var totalAmount float64
	var orderItems []entity.OrderItem

	for _, item := range cart.Items {
		if item.Product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + item.Product.Name)
		}

		totalAmount += item.Product.Price * float64(item.Quantity)
		orderItems = append(orderItems, entity.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})

		item.Product.Stock -= item.Quantity
		s.productRepo.Update(ctx, &item.Product)
	}

	order := &entity.Order{
		UserID:      userID,
		TotalAmount: totalAmount,
		OrderItems:  orderItems,
		Status:      "pending",
	}

	if err := s.orderRepo.Create(ctx, order); err != nil {
		return nil, err
	}

	_ = s.cartRepo.ClearCart(ctx, cart.ID)

	return order, nil
}

func (s *orderService) GetOrder(ctx context.Context, orderID uuid.UUID) (*entity.Order, error) {
	return s.orderRepo.FindByID(ctx, orderID)
}

func (s *orderService) GetUserOrders(ctx context.Context, userID uuid.UUID) ([]entity.Order, error) {
	return s.orderRepo.FindByUserID(ctx, userID)
}

func (s *orderService) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status string) error {
	return s.orderRepo.UpdateStatus(ctx, orderID, status)
}
