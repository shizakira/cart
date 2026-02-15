package dto

type DeleteItemInput struct {
	UserID int `json:"user_id"`
	SkuID  int `json:"sku_id"`
}
