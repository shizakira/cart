package in_memory_storage_test

import (
	"context"
	"testing"

	"github.com/shizakira/cart/internal/adapter/in_memory_storage"
	"github.com/shizakira/cart/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage_Find_NotExists(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()

	cart, err := s.Find(ctx, 999)
	require.NoError(t, err)
	assert.Equal(t, domain.Cart{}, cart)
}

func TestStorage_Find_Exists(t *testing.T) {
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
