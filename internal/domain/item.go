package domain

type Item struct {
	skuID int
	count uint16
}

func NewItem(skuID int, count uint16) Item {
	return Item{skuID: skuID, count: count}
}

func (i Item) SkuID() int {
	return i.skuID
}

func (i Item) Count() uint16 {
	return i.count
}
