package usecase

import (
	"context"

	"github.com/shizakira/cart/internal/domain"
	"github.com/shizakira/cart/internal/model"
)

//go:generate mockery

type Storage interface {
	AddItem(ctx context.Context, userID int, item domain.Item) error
	RemoveItem(ctx context.Context, userID int, skuID int) error
	Find(ctx context.Context, userID int) (domain.Cart, error)
	Clear(ctx context.Context, userID int) error
}

type ProductService interface {
	IsProductExist(ctx context.Context, skuID int) (bool, error)
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
