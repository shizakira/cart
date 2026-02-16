package in_memory_storage

import (
	"sync"

	"github.com/shizakira/cart/internal/domain"
)

type Storage struct {
	mu    sync.RWMutex
	carts map[int]map[int]domain.Item
}

func New() *Storage {
	return &Storage{
		carts: make(map[int]map[int]domain.Item),
	}
}
