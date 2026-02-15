package dto

type GetItemsInput struct {
	UserID int `json:"user_id"`
}

type GetItemsOutput struct {
	Items      []Item `json:"items"`
	TotalPrice uint   `json:"total_price"`
}

type Item struct {
	SkuID int    `json:"sku_id"`
	Name  string `json:"name"`
	Count uint   `json:"count"`
	Price uint   `json:"price"`
}
