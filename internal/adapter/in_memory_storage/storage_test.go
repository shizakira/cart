package in_memory_storage_test

import (
	"context"
	"testing"

	"github.com/shizakira/cart/internal/adapter/in_memory_storage"
	"github.com/shizakira/cart/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestStorage_Find_NotFound(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()
	userID := 1007

	cart, err := s.Find(ctx, userID)

	require.NoError(t, err)
	require.Nil(t, cart)
}

func TestStorage_Find_Found(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()
	userID := 1007
	expectedCart := domain.NewCart(userID)

	err := s.Save(ctx, expectedCart)
	require.NoError(t, err)

	cart, err := s.Find(ctx, userID)

	require.NoError(t, err)
	require.NotNil(t, cart)
	require.Equal(t, userID, cart.UserID())
}

func TestStorage_Save_UpdateExistingCart(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()
	userID := 1007
	cart := domain.NewCart(userID)
	item := domain.NewItem(2008, 1)
	cart.AddItem(item)

	err := s.Save(ctx, cart)
	require.NoError(t, err)

	cart, err = s.Find(ctx, userID)
	require.NoError(t, err)
	require.NotNil(t, cart)

	item = domain.NewItem(2008, 2)
	cart.AddItem(item)

	err = s.Save(ctx, cart)
	require.NoError(t, err)

	cart, err = s.Find(ctx, userID)
	require.NoError(t, err)
	require.Len(t, cart.Items(), 1)
	require.Equal(t, uint16(3), cart.Items()[0].Count())
}

func TestStorage_Save_ClearExistingCart(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()
	userID := 1007

	cart := domain.NewCart(userID)
	item := domain.NewItem(2008, 1)
	cart.AddItem(item)

	err := s.Save(ctx, cart)
	require.NoError(t, err)

	err = s.Clear(ctx, userID)
	require.NoError(t, err)

	cart, err = s.Find(ctx, userID)
	require.NoError(t, err)
	require.NotNil(t, cart)
}

func TestStorage_Save_ClearNotExistingCart(t *testing.T) {
	ctx := context.Background()
	s := in_memory_storage.New()
	userID := 1007

	err := s.Clear(ctx, userID)
	require.NoError(t, err)

	cart, err := s.Find(ctx, userID)
	require.NoError(t, err)
	require.Nil(t, cart)
}
