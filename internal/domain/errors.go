package domain

import "errors"

var (
	ErrCartIsEmpty       = errors.New("cart is empty")
	ErrCartNotFound      = errors.New("cart not found")
	ErrInsufficientStock = errors.New("insufficient stock")
)
