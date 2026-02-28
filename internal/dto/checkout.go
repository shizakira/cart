package dto

import val "github.com/go-ozzo/ozzo-validation"

type CheckoutInput struct {
	UserID int
}

func (d CheckoutInput) Validate() error {
	return val.ValidateStruct(&d,
		val.Field(&d.UserID, val.Required, val.Min(1).Error("user_id must be >= 1")),
	)
}

type CheckoutOutput struct {
	OrderID int64 `json:"order_id"`
}
