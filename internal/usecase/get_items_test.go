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
	"github.com/stretchr/testify/require"
)

func TestCart_GetItems_Success(t *testing.T) {
	ctx := context.Background()
	input := dto.GetItemsInput{UserID: 1007}

	cart := domain.Cart{
		Items: []domain.Item{
			{SkuID: 2008, Count: 2},
			{SkuID: 1001, Count: 3},
			{SkuID: 3002, Count: 1},
		},
	}
	expectedTotalSum := uint(2500)

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).Return(cart, nil).Once()

	productSvc := &mocks.ProductService{}
	productSvc.On("GetProduct", ctx, 2008).Return(model.Product{Name: "Product A", Price: 500}, nil).Once()
	productSvc.On("GetProduct", ctx, 1001).Return(model.Product{Name: "Product B", Price: 300}, nil).Once()
	productSvc.On("GetProduct", ctx, 3002).Return(model.Product{Name: "Product C", Price: 600}, nil).Once()

	uc := usecase.NewCart(storage, productSvc)

	output, err := uc.GetItems(ctx, input)

	require.NoError(t, err)
	require.Len(t, output.Items, 3)
	require.Equal(t, expectedTotalSum, output.TotalPrice)

	require.Equal(t, 1001, output.Items[0].SkuID)
	require.Equal(t, 2008, output.Items[1].SkuID)
	require.Equal(t, 3002, output.Items[2].SkuID)

	productSvc.AssertExpectations(t)
}

func TestCart_GetItems_CartNotFound(t *testing.T) {
	ctx := context.Background()
	input := dto.GetItemsInput{UserID: 1007}

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).Return(domain.Cart{}, domain.ErrCartNotFound).Once()

	uc := usecase.NewCart(storage, nil)

	_, err := uc.GetItems(ctx, input)

	require.Error(t, err)
	require.ErrorIs(t, err, domain.ErrCartNotFound)

	storage.AssertExpectations(t)
}

func TestCart_GetItems_CartEmpty(t *testing.T) {
	ctx := context.Background()
	input := dto.GetItemsInput{UserID: 1007}

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).Return(domain.Cart{Items: nil}, nil).Once()

	uc := usecase.NewCart(storage, nil)

	_, err := uc.GetItems(ctx, input)

	require.Error(t, err)
	require.ErrorIs(t, err, domain.ErrCartIsEmpty)

	storage.AssertExpectations(t)
}

func TestCart_GetItems_ProductServiceError(t *testing.T) {
	ctx := context.Background()
	input := dto.GetItemsInput{UserID: 1007}

	cart := domain.Cart{
		Items: []domain.Item{{SkuID: 2008, Count: 1}},
	}

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).Return(cart, nil).Once()

	productSvc := &mocks.ProductService{}
	productSvc.On("GetProduct", ctx, 2008).Return(model.Product{}, assert.AnError).Once()

	uc := usecase.NewCart(storage, productSvc)

	_, err := uc.GetItems(ctx, input)

	require.Error(t, err)

	storage.AssertExpectations(t)
}
