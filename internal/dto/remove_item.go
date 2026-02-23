package dto

import val "github.com/go-ozzo/ozzo-validation"

type RemoveItemInput struct {
	UserID int `json:"user_id"`
	SkuID  int `json:"sku_id"`
}

func (d RemoveItemInput) Validate() error {
	return val.ValidateStruct(&d,
		val.Field(&d.UserID, val.Required, val.Min(1).Error("sku_id must be >= 1")),
		val.Field(&d.SkuID, val.Required, val.Min(1).Error("sku_id must be >= 1")),
	)
}
