package repository

import (
	"context"
	"ecommerce/internal/domain/entity"
	domrepo "ecommerce/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domrepo.ProductRepository {
	return &productRepositoryImpl{db: db}
}

func (r *productRepositoryImpl) Create(ctx context.Context, product *entity.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	var product entity.Product
	if err := r.db.WithContext(ctx).Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepositoryImpl) FindAll(ctx context.Context, page, limit int) ([]entity.Product, int64, error) {
	var products []entity.Product
	var total int64
	offset := (page - 1) * limit

	r.db.Model(&entity.Product{}).Count(&total)
	err := r.db.WithContext(ctx).Preload("Category").Limit(limit).Offset(offset).Find(&products).Error
	return products, total, err
}

func (r *productRepositoryImpl) Update(ctx context.Context, product *entity.Product) error {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *productRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&entity.Product{}, id).Error
}

func (r *productRepositoryImpl) Search(ctx context.Context, query string, page, limit int) ([]entity.Product, int64, error) {
	var products []entity.Product
	var total int64
	offset := (page - 1) * limit

	searchQuery := r.db.Model(&entity.Product{}).Where("name ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%")
	searchQuery.Count(&total)
	err := searchQuery.Preload("Category").Limit(limit).Offset(offset).Find(&products).Error
	return products, total, err
}
