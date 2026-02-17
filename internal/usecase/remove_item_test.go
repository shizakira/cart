package usecase_test

import (
	"context"
	"testing"

	"github.com/shizakira/cart/internal/dto"
	"github.com/shizakira/cart/internal/usecase"
	"github.com/shizakira/cart/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCart_RemoveItem_Success(t *testing.T) {
	ctx := context.Background()
	input := dto.RemoveItemInput{
		UserID: 1007,
		SkuID:  2008,
	}

	storage := &mocks.Storage{}
	storage.On("RemoveItem", ctx, input.UserID, input.SkuID).
		Return(nil).
		Once()

	uc := usecase.NewCart(storage, nil)

	err := uc.RemoveItem(ctx, input)

	require.NoError(t, err)
	storage.AssertExpectations(t)
}

func TestCart_RemoveItem_StorageError(t *testing.T) {
	ctx := context.Background()
	input := dto.RemoveItemInput{
		UserID: 1007,
		SkuID:  2008,
	}

	storage := &mocks.Storage{}
	storage.On("RemoveItem", ctx, input.UserID, input.SkuID).
		Return(assert.AnError).
		Once()

	uc := usecase.NewCart(storage, nil)

	err := uc.RemoveItem(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "storage.RemoveItem")

	storage.AssertExpectations(t)
}
