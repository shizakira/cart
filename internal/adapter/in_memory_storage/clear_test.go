package in_memory_storage_test

import (
	"context"
	"testing"

	"github.com/shizakira/cart/internal/adapter/in_memory_storage"
	"github.com/shizakira/cart/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage_Clear_Exists(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()

	err := s.AddItem(ctx, 1, domain.Item{SkuID: 1001, Count: 1})
	require.NoError(t, err)

	err = s.Clear(ctx, 1)
	require.NoError(t, err)

	cart, err := s.Find(ctx, 1)

	require.NoError(t, err)
	assert.Equal(t, domain.Cart{}, cart)
}

func TestStorage_Clear_NotExistsNoOp(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()

	err := s.Clear(ctx, 999)
	require.NoError(t, err)

	cart, err := s.Find(ctx, 999)

	require.NoError(t, err)
	assert.Equal(t, domain.Cart{}, cart)
}
