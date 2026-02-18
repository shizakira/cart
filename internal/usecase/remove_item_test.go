package usecase_test

import (
	"context"
	"testing"

	"github.com/shizakira/cart/internal/domain"
	"github.com/shizakira/cart/internal/dto"
	"github.com/shizakira/cart/internal/usecase"
	"github.com/shizakira/cart/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCart_RemoveItem_Success_CartNotFound(t *testing.T) {
	ctx := context.Background()
	input := dto.RemoveItemInput{
		UserID: 1007,
		SkuID:  2008,
	}

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return(nil, nil).
		Once()
	storage.AssertNotCalled(t, "Save")

	uc := usecase.NewCart(storage, nil)

	err := uc.RemoveItem(ctx, input)

	require.NoError(t, err)
	storage.AssertExpectations(t)
}

func TestCart_RemoveItem_Success_ItemRemoved(t *testing.T) {
	ctx := context.Background()
	input := dto.RemoveItemInput{
		UserID: 1007,
		SkuID:  2008,
	}

	existingCart := domain.NewCart(input.UserID)
	existingItem := domain.NewItem(input.SkuID, 3)
	existingCart.AddItem(existingItem)

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return(existingCart, nil).
		Once()
	storage.On("Save", ctx, mock.MatchedBy(func(cart *domain.Cart) bool {
		return cart.UserID() == input.UserID &&
			len(cart.Items()) == 0
	})).
		Return(nil).
		Once()

	uc := usecase.NewCart(storage, nil)

	err := uc.RemoveItem(ctx, input)

	require.NoError(t, err)
	storage.AssertExpectations(t)
}

func TestCart_RemoveItem_Success_ItemNotInCart(t *testing.T) {
	ctx := context.Background()
	input := dto.RemoveItemInput{
		UserID: 1007,
		SkuID:  2008,
	}

	existingCart := domain.NewCart(input.UserID)

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return(existingCart, nil).
		Once()
	storage.On("Save", ctx, mock.MatchedBy(func(cart *domain.Cart) bool {
		return cart.UserID() == input.UserID &&
			len(cart.Items()) == 0
	})).
		Return(nil).
		Once()

	uc := usecase.NewCart(storage, nil)

	err := uc.RemoveItem(ctx, input)

	require.NoError(t, err)
	storage.AssertExpectations(t)
}

func TestCart_RemoveItem_StorageFindReturnsError(t *testing.T) {
	ctx := context.Background()
	input := dto.RemoveItemInput{
		UserID: 1007,
		SkuID:  2008,
	}

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return(nil, assert.AnError).
		Once()
	storage.AssertNotCalled(t, "Save")

	uc := usecase.NewCart(storage, nil)

	err := uc.RemoveItem(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "storage.Find")
	storage.AssertExpectations(t)
}

func TestCart_RemoveItem_StorageSaveReturnsError(t *testing.T) {
	ctx := context.Background()
	input := dto.RemoveItemInput{
		UserID: 1007,
		SkuID:  2008,
	}

	existingCart := domain.NewCart(input.UserID)
	existingItem := domain.NewItem(input.SkuID, 3)
	existingCart.AddItem(existingItem)

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return(existingCart, nil).
		Once()
	storage.On("Save", ctx, mock.MatchedBy(func(cart *domain.Cart) bool {
		return cart.UserID() == input.UserID &&
			len(cart.Items()) == 0
	})).
		Return(assert.AnError).
		Once()

	uc := usecase.NewCart(storage, nil)

	err := uc.RemoveItem(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "storage.AddItem")
	storage.AssertExpectations(t)
}
