package domain

type Item struct {
	SkuID int
	Count uint
}

func NewItem(skuID int, count uint) Item {
	return Item{SkuID: skuID, Count: count}
}
