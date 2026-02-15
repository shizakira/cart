package usecase

import (
	"context"
	"fmt"
	"slices"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/shizakira/cart/internal/domain"
	"github.com/shizakira/cart/internal/dto"
)

func (c *Cart) GetItems(ctx context.Context, input dto.GetItemsInput) (dto.GetItemsOutput, error) {
	var output dto.GetItemsOutput

	cart, err := c.storage.Find(ctx, input.UserID)
	if err != nil {
		return output, fmt.Errorf("storage.Find: %w", err)
	}
	if len(cart.Items) == 0 {
		return output, domain.ErrCartIsEmpty
	}

	var totalPrice uint
	items := make([]dto.Item, 0, len(cart.Items))

	var g errgroup.Group
	mx := sync.Mutex{}

	for _, item := range cart.Items {
		g.Go(func() error {
			product, err := c.productService.GetProduct(ctx, item.SkuID)
			if err != nil {
				return fmt.Errorf("productService.GetProduct: %w", err)
			}

			mx.Lock()

			totalPrice += product.Price * item.Count
			items = append(items, dto.Item{
				SkuID: item.SkuID,
				Name:  product.Name,
				Count: item.Count,
				Price: product.Price,
			})

			mx.Unlock()

			return nil
		})

	}

	if err = g.Wait(); err != nil {
		return output, fmt.Errorf("g.Wait: %w", err)
	}

	slices.SortFunc(items, func(item dto.Item, item2 dto.Item) int {
		return item.SkuID - item2.SkuID
	})

	output.Items = items
	output.TotalPrice = totalPrice

	return output, nil
}
