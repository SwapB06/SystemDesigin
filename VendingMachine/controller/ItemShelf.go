package controller

type ItemShelf struct {
	Code    int
	Item    Item
	SoldOut bool
}

func NewItemShelf() *ItemShelf {
	return &ItemShelf{}
}

func (is ItemShelf) GetCode() int {
	return is.Code
}
func (is *ItemShelf) SetCode(code int) {
	is.Code = code
}
func (is ItemShelf) GetItem() Item {
	return is.Item
}
func (is *ItemShelf) SetItem(item *Item) {
	is.Item = *item
}
func (is ItemShelf) IsSoldOut() bool {
	return is.SoldOut
}
func (is *ItemShelf) SetSoldOut(isit bool) {
	is.SoldOut = isit
}
