package controller

import "fmt"

type Inventory struct {
	Inventory []*ItemShelf
}

func NewInventory(itemCount int) Inventory {
	inv := Inventory{
		Inventory: make([]*ItemShelf, itemCount),
	}
	inv.initialEmptyInventory()
	return inv
}

func (i Inventory) GetInventory() *[]*ItemShelf {
	return &i.Inventory
}

func (I *Inventory) initialEmptyInventory() {
	startCode := 101
	for i := 0; i < len(I.Inventory); i++ {
		space := NewItemShelf()
		space.SetCode(startCode)
		space.SetSoldOut(true)
		I.Inventory[i] = space
		startCode++
	}
}

func (i *Inventory) AddItem(item Item, CodeNumber int) {
	for _, itemShelf := range i.Inventory {
		if itemShelf.Code == CodeNumber {
			if itemShelf.IsSoldOut() {
				itemShelf.Item = item
				itemShelf.SetSoldOut(false)
			} else {
				fmt.Println("ERROR: Already item is present, you cannot add item here")
			}
		}
	}
}

func (i Inventory) GetItem(codeNumber int) Item {
	for _, itemShelf := range i.Inventory {
		if itemShelf.Code == codeNumber {
			if itemShelf.IsSoldOut() {
				fmt.Println("ERROR: Item Already Sold out")
			} else {
				return itemShelf.Item
			}
		}
	}
	fmt.Println("ERROR: Invalid Code")
	return Item{}
}

func (i *Inventory) UpdateSoldOutItem(codeNumber int) {
	for _, itemShelf := range i.Inventory {
		if itemShelf.Code == codeNumber {
			itemShelf.SetSoldOut(true)
		}
	}
}
