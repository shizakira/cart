package dto

import (
	val "github.com/go-ozzo/ozzo-validation"
)

type AddItemInput struct {
	UserID int
	SkuID  int
	Count  int
}

func (d AddItemInput) Validate() error {
	return val.ValidateStruct(&d,
		val.Field(&d.UserID, val.Required, val.Min(1).Error("sku_id must be >= 1")),
		val.Field(&d.SkuID, val.Required, val.Min(1).Error("sku_id must be >= 1")),
		val.Field(&d.Count, val.Required, val.Min(1).Error("count must be >= 1")),
	)
}
