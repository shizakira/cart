package usecase

import (
	"context"

	"github.com/shizakira/cart/internal/domain"
)

//go:generate mockery

type Storage interface {
	AddItem(ctx context.Context, userID int, skuID int, count uint) error
	RemoveItem(ctx context.Context, userID int, skuID int) error
	Find(ctx context.Context, userID int) (domain.Cart, error)
	Clear(ctx context.Context, userID int) error
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
