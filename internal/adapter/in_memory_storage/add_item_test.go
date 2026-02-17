package in_memory_storage_test

import (
	"context"
	"testing"

	"github.com/shizakira/cart/internal/adapter/in_memory_storage"
	"github.com/shizakira/cart/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage_AddItem_NewUserNewItem(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()

	item := domain.Item{SkuID: 1001, Count: 3}
	err := s.AddItem(ctx, 1, item)
	require.NoError(t, err)

	cart, err := s.Find(ctx, 1)

	require.NoError(t, err)
	assert.Equal(t, 1, cart.UserID)
	assert.Len(t, cart.Items, 1)
	assert.Equal(t, item, cart.Items[0])
}

func TestStorage_AddItem_IncrementExistingItem(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()

	err := s.AddItem(ctx, 1, domain.Item{SkuID: 1001, Count: 3})
	require.NoError(t, err)
	err = s.AddItem(ctx, 1, domain.Item{SkuID: 1001, Count: 5})
	require.NoError(t, err)

	cart, err := s.Find(ctx, 1)

	require.NoError(t, err)
	assert.Equal(t, 1, cart.UserID)
	assert.Len(t, cart.Items, 1)
	assert.Equal(t, uint(8), cart.Items[0].Count)
}

func TestStorage_AddItem_NewItemInExistingCart(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()

	err := s.AddItem(ctx, 1, domain.Item{SkuID: 1001, Count: 1})
	require.NoError(t, err)
	err = s.AddItem(ctx, 1, domain.Item{SkuID: 1002, Count: 2})
	require.NoError(t, err)

	cart, err := s.Find(ctx, 1)

	require.NoError(t, err)
	assert.Equal(t, 1, cart.UserID)
	assert.Len(t, cart.Items, 2)

	expected := []domain.Item{
		{SkuID: 1001, Count: 1},
		{SkuID: 1002, Count: 2},
	}
	assert.ElementsMatch(t, expected, cart.Items)
}
