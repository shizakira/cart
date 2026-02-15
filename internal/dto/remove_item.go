package dto

type RemoveItemInput struct {
	UserID int `json:"user_id"`
	SkuID  int `json:"sku_id"`
}
