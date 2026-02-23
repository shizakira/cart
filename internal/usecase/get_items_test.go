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
	input := dto.GetItemsInput{
		UserID: 1007,
	}

	cart := domain.NewCart(input.UserID)
	cart.AddItem(domain.NewItem(2008, 2))
	cart.AddItem(domain.NewItem(1001, 3))
	cart.AddItem(domain.NewItem(3002, 1))

	expectedTotalPrice := uint32(2500)

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return(cart, nil).
		Once()

	productSvc := &mocks.ProductService{}
	productSvc.On("GetProduct", ctx, 1001).
		Return(model.Product{Name: "Product B", Price: 300}, nil).
		Once()
	productSvc.On("GetProduct", ctx, 2008).
		Return(model.Product{Name: "Product A", Price: 500}, nil).
		Once()
	productSvc.On("GetProduct", ctx, 3002).
		Return(model.Product{Name: "Product C", Price: 600}, nil).
		Once()

	uc := usecase.NewCart(storage, productSvc)

	output, err := uc.GetItems(ctx, input)

	require.NoError(t, err)
	require.Len(t, output.Items, 3)
	require.Equal(t, expectedTotalPrice, output.TotalPrice)

	require.Equal(t, 1001, output.Items[0].SkuID)
	require.Equal(t, "Product B", output.Items[0].Name)
	require.Equal(t, uint16(3), output.Items[0].Count)
	require.Equal(t, uint32(300), output.Items[0].Price)

	require.Equal(t, 2008, output.Items[1].SkuID)
	require.Equal(t, "Product A", output.Items[1].Name)
	require.Equal(t, uint16(2), output.Items[1].Count)
	require.Equal(t, uint32(500), output.Items[1].Price)

	require.Equal(t, 3002, output.Items[2].SkuID)
	require.Equal(t, "Product C", output.Items[2].Name)
	require.Equal(t, uint16(1), output.Items[2].Count)
	require.Equal(t, uint32(600), output.Items[2].Price)

	storage.AssertExpectations(t)
	productSvc.AssertExpectations(t)
}

func TestCart_GetItems_StorageFindReturnsError(t *testing.T) {
	ctx := context.Background()
	input := dto.GetItemsInput{
		UserID: 1007,
	}

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return((*domain.Cart)(nil), assert.AnError).
		Once()

	uc := usecase.NewCart(storage, nil)

	_, err := uc.GetItems(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "storage.Find")

	storage.AssertExpectations(t)
}

func TestCart_GetItems_CartEmpty(t *testing.T) {
	ctx := context.Background()
	input := dto.GetItemsInput{
		UserID: 1007,
	}

	cart := domain.NewCart(input.UserID)

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return(cart, nil).
		Once()

	uc := usecase.NewCart(storage, nil)

	_, err := uc.GetItems(ctx, input)

	require.ErrorIs(t, err, domain.ErrCartIsEmpty)

	storage.AssertExpectations(t)
}

func TestCart_GetItems_ProductServiceReturnsError(t *testing.T) {
	ctx := context.Background()
	input := dto.GetItemsInput{
		UserID: 1007,
	}

	cart := domain.NewCart(input.UserID)
	cart.AddItem(domain.NewItem(2008, 1))

	storage := &mocks.Storage{}
	storage.On("Find", ctx, input.UserID).
		Return(cart, nil).
		Once()

	productSvc := &mocks.ProductService{}
	productSvc.On("GetProduct", ctx, 2008).
		Return(model.Product{}, assert.AnError).
		Once()

	uc := usecase.NewCart(storage, productSvc)

	_, err := uc.GetItems(ctx, input)

	require.Error(t, err)
	require.Contains(t, err.Error(), "productService.GetProduct")

	storage.AssertExpectations(t)
	productSvc.AssertExpectations(t)
}
