package handlers

import (
	"github.com/shizakira/cart/internal/usecase"
)

type Handlers struct {
	usecase *usecase.Cart
}

func NewHandlers(usecase *usecase.Cart) *Handlers {
	return &Handlers{usecase: usecase}
}
