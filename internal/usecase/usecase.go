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
type LomsService interface {
	CreateOrder(ctx context.Context, userID int, items []domain.Item) (int64, error)
	StocksInfo(ctx context.Context, sku uint32) (uint64, error)
}
type Cart struct {
	storage        Storage
	productService ProductService
	lomsService    LomsService
}

func NewCart(storage Storage, productService ProductService, lomsService LomsService) *Cart {
	return &Cart{
		storage:        storage,
		productService: productService,
		lomsService:    lomsService,
	}
}
