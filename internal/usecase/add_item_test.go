package usecase_test

import (
	"context"
	"testing"

	"github.com/shizakira/cart/internal/domain"
	"github.com/shizakira/cart/internal/dto"
	"github.com/shizakira/cart/internal/model"
	"github.com/shizakira/cart/internal/usecase"
	"github.com/shizakira/cart/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCart_AddItem_Success_NewItem(t *testing.T) {
	ctx := context.Background()
	input := dto.AddItemInput{
		UserID: 1007,
		SkuID:  2008,
		Count:  3,
	}
	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return(nil, nil).
		Once()
	storage.On("Save", ctx, mock.MatchedBy(func(cart *domain.Cart) bool {
		return cart.UserID() == input.UserID &&
			len(cart.Items()) == 1 &&
			cart.Items()[0].SkuID() == input.SkuID &&
			cart.Items()[0].Count() == uint16(input.Count)
	})).
		Return(nil).
		Once()

	productSvc := &mocks.ProductService{}
	productSvc.On("GetProduct", ctx, input.SkuID).
		Return(model.Product{}, nil).
		Once()

	lomsSvc := &mocks.LomsService{}
	lomsSvc.On("StocksInfo", ctx, uint32(input.SkuID)).
		Return(uint64(100), nil).
		Once()

	uc := usecase.NewCart(storage, productSvc, lomsSvc)

	err := uc.AddItem(ctx, input)

	require.NoError(t, err)
	storage.AssertExpectations(t)
	productSvc.AssertExpectations(t)
	lomsSvc.AssertExpectations(t)
}

func TestCart_AddItem_InsufficientStock(t *testing.T) {
	ctx := context.Background()
	input := dto.AddItemInput{
		UserID: 1007,
		SkuID:  2008,
		Count:  100,
	}

	storage := &mocks.Storage{}
	storage.AssertNotCalled(t, "Find")
	storage.AssertNotCalled(t, "Save")

	productSvc := &mocks.ProductService{}
	productSvc.On("GetProduct", ctx, input.SkuID).
		Return(model.Product{}, nil).
		Once()

	lomsSvc := &mocks.LomsService{}
	lomsSvc.On("StocksInfo", ctx, uint32(input.SkuID)).
		Return(uint64(5), nil).
		Once()

	uc := usecase.NewCart(storage, productSvc, lomsSvc)

	err := uc.AddItem(ctx, input)

	require.Error(t, err)
	require.ErrorIs(t, err, domain.ErrInsufficientStock)
	storage.AssertExpectations(t)
	productSvc.AssertExpectations(t)
	lomsSvc.AssertExpectations(t)
}

func TestCart_AddItem_LomsServiceReturnsError(t *testing.T) {
	ctx := context.Background()
	input := dto.AddItemInput{
		UserID: 1007,
		SkuID:  2008,
		Count:  1,
	}

	storage := &mocks.Storage{}
	storage.AssertNotCalled(t, "Find")
	storage.AssertNotCalled(t, "Save")

	productSvc := &mocks.ProductService{}
	productSvc.On("GetProduct", ctx, input.SkuID).
		Return(model.Product{}, nil).
		Once()

	lomsSvc := &mocks.LomsService{}
	lomsSvc.On("StocksInfo", ctx, uint32(input.SkuID)).
		Return(uint64(0), assert.AnError).
		Once()

	uc := usecase.NewCart(storage, productSvc, lomsSvc)

	err := uc.AddItem(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "lomsService.StocksInfo")
	storage.AssertExpectations(t)
	productSvc.AssertExpectations(t)
	lomsSvc.AssertExpectations(t)
}

func TestCart_AddItem_ProductNotFound(t *testing.T) {
	ctx := context.Background()
	input := dto.AddItemInput{
		UserID: 1007,
		SkuID:  9999,
		Count:  1,
	}

	storage := &mocks.Storage{}
	storage.AssertNotCalled(t, "Find")
	storage.AssertNotCalled(t, "Save")

	productSvc := &mocks.ProductService{}
	productSvc.On("GetProduct", ctx, input.SkuID).
		Return(model.Product{}, model.ErrProductNotFound).
		Once()

	lomsSvc := &mocks.LomsService{}
	lomsSvc.AssertNotCalled(t, "StocksInfo")

	uc := usecase.NewCart(storage, productSvc, lomsSvc)

	err := uc.AddItem(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "productService.GetProduct")
	require.ErrorIs(t, err, model.ErrProductNotFound)
	productSvc.AssertExpectations(t)
	storage.AssertExpectations(t)
	lomsSvc.AssertExpectations(t)
}

func TestCart_AddItem_ProductServiceReturnsError(t *testing.T) {
	ctx := context.Background()
	input := dto.AddItemInput{
		UserID: 1007,
		SkuID:  2008,
		Count:  1,
	}

	storage := &mocks.Storage{}
	storage.AssertNotCalled(t, "Find")
	storage.AssertNotCalled(t, "Save")

	productSvc := &mocks.ProductService{}
	productSvc.On("GetProduct", ctx, input.SkuID).
		Return(model.Product{}, assert.AnError).
		Once()

	lomsSvc := &mocks.LomsService{}
	lomsSvc.AssertNotCalled(t, "StocksInfo")

	uc := usecase.NewCart(storage, productSvc, lomsSvc)

	err := uc.AddItem(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "productService.GetProduct")
	productSvc.AssertExpectations(t)
	storage.AssertExpectations(t)
	lomsSvc.AssertExpectations(t)
}

func TestCart_AddItem_StorageFindReturnsError(t *testing.T) {
	ctx := context.Background()
	input := dto.AddItemInput{
		UserID: 1007,
		SkuID:  2008,
		Count:  4,
	}

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return(nil, assert.AnError).
		Once()
	storage.AssertNotCalled(t, "Save")

	productSvc := &mocks.ProductService{}
	productSvc.On("GetProduct", ctx, input.SkuID).
		Return(model.Product{}, nil).
		Once()

	lomsSvc := &mocks.LomsService{}
	lomsSvc.On("StocksInfo", ctx, uint32(input.SkuID)).
		Return(uint64(100), nil).
		Once()

	uc := usecase.NewCart(storage, productSvc, lomsSvc)

	err := uc.AddItem(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "storage.Find")
	storage.AssertExpectations(t)
	productSvc.AssertExpectations(t)
	lomsSvc.AssertExpectations(t)
}

func TestCart_AddItem_StorageSaveReturnsError(t *testing.T) {
	ctx := context.Background()
	input := dto.AddItemInput{
		UserID: 1007,
		SkuID:  2008,
		Count:  4,
	}

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return(nil, nil).
		Once()
	storage.On("Save", ctx, mock.MatchedBy(func(cart *domain.Cart) bool {
		return cart.UserID() == input.UserID &&
			len(cart.Items()) == 1 &&
			cart.Items()[0].SkuID() == input.SkuID &&
			cart.Items()[0].Count() == uint16(input.Count)
	})).
		Return(assert.AnError).
		Once()

	productSvc := &mocks.ProductService{}
	productSvc.On("GetProduct", ctx, input.SkuID).
		Return(model.Product{}, nil).
		Once()

	lomsSvc := &mocks.LomsService{}
	lomsSvc.On("StocksInfo", ctx, uint32(input.SkuID)).
		Return(uint64(100), nil).
		Once()

	uc := usecase.NewCart(storage, productSvc, lomsSvc)

	err := uc.AddItem(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "storage.Save")
	storage.AssertExpectations(t)
	productSvc.AssertExpectations(t)
	lomsSvc.AssertExpectations(t)
}
