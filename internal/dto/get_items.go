package dto

import val "github.com/go-ozzo/ozzo-validation"

type GetItemsInput struct {
	UserID int
}

func (d GetItemsInput) Validate() error {
	return val.ValidateStruct(&d,
		val.Field(&d.UserID, val.Required, val.Min(1).Error("sku_id must be >= 1")),
	)
}

type GetItemsOutput struct {
	Items      []Item `json:"items"`
	TotalPrice uint32 `json:"total_price"`
}

type Item struct {
	SkuID int    `json:"sku_id"`
	Name  string `json:"name"`
	Count uint16 `json:"count"`
	Price uint32 `json:"price"`
}
