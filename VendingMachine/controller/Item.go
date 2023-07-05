package controller

type Item struct {
	Type  ItemType
	Price int
}

func NewItem() *Item {
	return &Item{}
}
func (i Item) GetType() ItemType {
	return i.Type
}

func (i *Item) SetType(typ ItemType) {
	i.Type = typ
}

func (i Item) GetPrice() int {
	return i.Price
}

func (i *Item) SetPrice(price int) {
	i.Price = price
}
