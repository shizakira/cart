package domain

type Cart struct {
	userID int
	items  map[int]Item
}

func NewCart(userID int) *Cart {
	return &Cart{userID: userID, items: make(map[int]Item)}
}

func (c *Cart) AddItem(item Item) {
	existItem, ok := c.items[item.skuID]
	if !ok {
		c.items[item.skuID] = item

		return
	}

	existItem.count += item.count
	c.items[item.skuID] = existItem
}

func (c *Cart) RemoveItem(skuID int) {
	delete(c.items, skuID)
}

func (c *Cart) IsEmpty() bool {
	return len(c.items) == 0
}

func (c *Cart) UserID() int {
	return c.userID
}

func (c *Cart) Items() []Item {
	items := make([]Item, 0, len(c.items))
	for _, v := range c.items {
		items = append(items, v)
	}

	return items
}
