package usecase_test

import (
	"context"
	"testing"

	"github.com/shizakira/cart/internal/domain"
	"github.com/shizakira/cart/internal/dto"
	"github.com/shizakira/cart/internal/usecase"
	"github.com/shizakira/cart/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
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
	storage.On("AddItem", ctx, input.UserID, input.SkuID, input.Count).
		Return(nil).
		Once()

	productSvc := &mocks.ProductService{}
	productSvc.On("IsProductExist", ctx, input.SkuID).
		Return(true, nil).
		Once()

	uc := usecase.NewCart(storage, productSvc)

	err := uc.AddItem(ctx, input)

	require.NoError(t, err)
	storage.AssertExpectations(t)
	productSvc.AssertExpectations(t)
}

func TestCart_AddItem_ProductNotFound(t *testing.T) {
	ctx := context.Background()
	input := dto.AddItemInput{
		UserID: 1007,
		SkuID:  9999,
		Count:  1,
	}

	storage := &mocks.Storage{}
	storage.AssertNotCalled(t, "AddItem")

	productSvc := &mocks.ProductService{}
	productSvc.On("IsProductExist", ctx, input.SkuID).
		Return(false, nil).
		Once()

	uc := usecase.NewCart(storage, productSvc)

	err := uc.AddItem(ctx, input)

	require.ErrorIs(t, err, domain.ErrItemNotFound)
	productSvc.AssertExpectations(t)
	storage.AssertExpectations(t)
}

func TestCart_AddItem_ProductServiceReturnsError(t *testing.T) {
	ctx := context.Background()
	input := dto.AddItemInput{
		UserID: 1007,
		SkuID:  2008,
		Count:  1,
	}

	storage := &mocks.Storage{}
	storage.AssertNotCalled(t, "AddItem")

	productSvc := &mocks.ProductService{}
	productSvc.On("IsProductExist", ctx, input.SkuID).
		Return(false, assert.AnError).
		Once()

	uc := usecase.NewCart(storage, productSvc)

	err := uc.AddItem(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "productService.IsProductExist")

	productSvc.AssertExpectations(t)
	storage.AssertExpectations(t)
}

func TestCart_AddItem_StorageReturnsError(t *testing.T) {
	ctx := context.Background()
	input := dto.AddItemInput{
		UserID: 1007,
		SkuID:  2008,
		Count:  4,
	}

	storage := &mocks.Storage{}
	storage.On("AddItem", ctx, input.UserID, input.SkuID, input.Count).
		Return(assert.AnError).
		Once()

	productSvc := &mocks.ProductService{}
	productSvc.On("IsProductExist", ctx, input.SkuID).
		Return(true, nil).
		Once()

	uc := usecase.NewCart(storage, productSvc)

	err := uc.AddItem(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "storage.AddItem")

	storage.AssertExpectations(t)
	productSvc.AssertExpectations(t)
}
