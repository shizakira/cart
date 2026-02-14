package usecase

import (
	"context"

	"github.com/shizakira/cart/internal/domain"
)

type Storage interface {
	FindByUserID(ctx context.Context, userID int) (domain.Cart, error)
	Create(ctx context.Context, userID int) (domain.Cart, error)
	Save(ctx context.Context, cart domain.Cart) error
	Delete(ctx context.Context)
}

type ProductService interface {
	IsProductExist(ctx context.Context, sku int) (bool, error)
}

type Cart struct {
	storage        Storage
	productService ProductService
}

func NewCart(storage Storage, productService ProductService) *Cart {
	return &Cart{
		storage:        storage,
		productService: productService,
	}
}
