package in_memory_storage_test

import (
	"context"
	"testing"

	"github.com/shizakira/cart/internal/adapter/in_memory_storage"
	"github.com/shizakira/cart/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage_RemoveItem_NotExistsNoOp(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()

	err := s.RemoveItem(ctx, 1, 1001)
	require.NoError(t, err)

	cart, err := s.Find(ctx, 1)
	require.NoError(t, err)
	assert.Equal(t, domain.Cart{}, cart)
}

func TestStorage_RemoveItem_RemoveOne(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()

	err := s.AddItem(ctx, 1, domain.Item{SkuID: 1001, Count: 1})
	require.NoError(t, err)
	err = s.AddItem(ctx, 1, domain.Item{SkuID: 1002, Count: 2})
	require.NoError(t, err)

	err = s.RemoveItem(ctx, 1, 1001)
	require.NoError(t, err)

	cart, err := s.Find(ctx, 1)

	require.NoError(t, err)
	assert.Equal(t, 1, cart.UserID)
	assert.Len(t, cart.Items, 1)
	assert.Equal(t, domain.Item{SkuID: 1002, Count: 2}, cart.Items[0])
}

func TestStorage_RemoveItem_RemoveLastDeletesCart(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()

	err := s.AddItem(ctx, 1, domain.Item{SkuID: 1001, Count: 1})
	require.NoError(t, err)

	err = s.RemoveItem(ctx, 1, 1001)
	require.NoError(t, err)

	cart, err := s.Find(ctx, 1)

	require.NoError(t, err)
	assert.Equal(t, domain.Cart{}, cart)
}

func TestStorage_RemoveItem_RemoveNonExistingInCartNoOp(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()

	err := s.AddItem(ctx, 1, domain.Item{SkuID: 1001, Count: 1})
	require.NoError(t, err)

	err = s.RemoveItem(ctx, 1, 1002)
	require.NoError(t, err)

	cart, err := s.Find(ctx, 1)

	require.NoError(t, err)
	assert.Equal(t, 1, cart.UserID)
	assert.Len(t, cart.Items, 1)
	assert.Equal(t, domain.Item{SkuID: 1001, Count: 1}, cart.Items[0])
}
