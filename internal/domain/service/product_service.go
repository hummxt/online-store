package service

import (
	"context"
	"ecommerce/internal/domain/entity"
	"ecommerce/internal/domain/repository"

	"github.com/google/uuid"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *entity.Product) error
	GetProduct(ctx context.Context, id uuid.UUID) (*entity.Product, error)
	ListProducts(ctx context.Context, page, limit int) ([]entity.Product, int64, error)
	UpdateProduct(ctx context.Context, product *entity.Product) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	SearchProducts(ctx context.Context, query string, page, limit int) ([]entity.Product, int64, error)

	CreateCategory(ctx context.Context, category *entity.Category) error
	GetCategory(ctx context.Context, id uuid.UUID) (*entity.Category, error)
	ListCategories(ctx context.Context) ([]entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id uuid.UUID) error
}

type productService struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductService(productRepo repository.ProductRepository, categoryRepo repository.CategoryRepository) ProductService {
	return &productService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, product *entity.Product) error {
	return s.productRepo.Create(ctx, product)
}

func (s *productService) GetProduct(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	return s.productRepo.FindByID(ctx, id)
}

func (s *productService) ListProducts(ctx context.Context, page, limit int) ([]entity.Product, int64, error) {
	return s.productRepo.FindAll(ctx, page, limit)
}

func (s *productService) UpdateProduct(ctx context.Context, product *entity.Product) error {
	return s.productRepo.Update(ctx, product)
}

func (s *productService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return s.productRepo.Delete(ctx, id)
}

func (s *productService) SearchProducts(ctx context.Context, query string, page, limit int) ([]entity.Product, int64, error) {
	return s.productRepo.Search(ctx, query, page, limit)
}

func (s *productService) CreateCategory(ctx context.Context, category *entity.Category) error {
	return s.categoryRepo.Create(ctx, category)
}

func (s *productService) GetCategory(ctx context.Context, id uuid.UUID) (*entity.Category, error) {
	return s.categoryRepo.FindByID(ctx, id)
}

func (s *productService) ListCategories(ctx context.Context) ([]entity.Category, error) {
	return s.categoryRepo.FindAll(ctx)
}

func (s *productService) UpdateCategory(ctx context.Context, category *entity.Category) error {
	return s.categoryRepo.Update(ctx, category)
}

func (s *productService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	return s.categoryRepo.Delete(ctx, id)
}
