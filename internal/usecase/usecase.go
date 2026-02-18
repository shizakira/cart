package usecase

import (
	"context"

	"github.com/shizakira/cart/internal/domain"
	"github.com/shizakira/cart/internal/model"
)

//go:generate mockery

type Storage interface {
	Find(ctx context.Context, userID int) (*domain.Cart, error)
	Save(ctx context.Context, cart *domain.Cart) error
	Clear(ctx context.Context, userID int) error
}

type ProductService interface {
	GetProduct(ctx context.Context, skuID int) (model.Product, error)
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
