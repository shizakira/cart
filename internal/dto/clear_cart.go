package dto

import val "github.com/go-ozzo/ozzo-validation"

type ClearCartInput struct {
	UserID int
}

func (d ClearCartInput) Validate() error {
	return val.ValidateStruct(&d,
		val.Field(&d.UserID, val.Required, val.Min(1).Error("sku_id must be >= 1")),
	)
}
