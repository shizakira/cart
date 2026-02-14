package domain

type Cart struct {
	UserID int
	Items  []Item
}

func NewCart(userID int, items []Item) Cart {
	return Cart{UserID: userID, Items: items}
}
