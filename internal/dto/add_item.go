package dto

type AddItemInput struct {
	UserID int  `json:"user_id"`
	SkuID  int  `json:"sku_id"`
	Count  uint `json:"count"`
}
