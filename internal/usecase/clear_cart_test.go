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

func TestCart_Clear_Success(t *testing.T) {
	ctx := context.Background()
	input := dto.ClearCartInput{UserID: 1007}

	storage := &mocks.Storage{}
	storage.On("Clear", ctx, input.UserID).
		Return(nil).
		Once()

	uc := usecase.NewCart(storage, nil, nil)

	err := uc.Clear(ctx, input)

	require.NoError(t, err)
	storage.AssertExpectations(t)
}

func TestCart_Clear_StorageError(t *testing.T) {
	ctx := context.Background()
	input := dto.ClearCartInput{UserID: 1007}

	storage := &mocks.Storage{}
	storage.On("Clear", ctx, input.UserID).
		Return(assert.AnError).
		Once()

	uc := usecase.NewCart(storage, nil, nil)

	err := uc.Clear(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "storage.Clear")

	storage.AssertExpectations(t)
}
