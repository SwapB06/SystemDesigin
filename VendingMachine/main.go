package main

import (
	"fmt"
	"vendingmachine/controller"
)

func main() {
	vm := controller.NewVendingMachine()
	fmt.Println("|")
	fmt.Println("Filling up the inventory")
	fmt.Println("|")

	FillUpInventory(&vm)
	DisplayInventory(vm)

	fmt.Println("|")
	fmt.Println("Clicking on insert coin button")
	fmt.Println("|")

	vendingState := vm.GetVendingMachineState()
	vendingState.ClickOnInsertCoinButton(&vm)

	vendingState = vm.GetVendingMachineState()
	vendingState.InsertCoin(&vm, controller.NICKEL)
	vendingState.InsertCoin(&vm, controller.QUARTER)

	fmt.Println("|")
	fmt.Println("Clicking on product selection button")
	fmt.Println("|")
	vendingState.ClickOnStartProductSelectionButton(&vm)

	vendingState = vm.GetVendingMachineState()
	vendingState.ChooseProduct(&vm, 102)

	DisplayInventory(vm)
}

func FillUpInventory(vendingMachine *controller.VendingMachine) {
	slots := vendingMachine.GetInventory()
	for i := 0; i < len(slots.Inventory); i++ {
		newItem := controller.NewItem()
		if i >= 0 && i < 3 {
			newItem.SetType(controller.COKE)
			newItem.SetPrice(12)
		} else if i >= 3 && i < 5 {
			newItem.SetType(controller.PEPSI)
			newItem.SetPrice(9)
		} else if i >= 5 && i < 7 {
			newItem.SetType(controller.JUICE)
			newItem.SetPrice(13)
		} else if i >= 7 && i < 10 {
			newItem.SetType(controller.SODA)
			newItem.SetPrice(7)
		}
		slots.Inventory[i].SetItem(newItem)
		slots.Inventory[i].SetSoldOut(false)
	}
}

func DisplayInventory(vm controller.VendingMachine) {
	slots := controller.VendingMachine.GetInventory(vm)
	for i := 0; i < len(slots.Inventory); i++ {
		fmt.Println("CodeNumber: ", slots.Inventory[i].GetCode(), " Item: ", slots.Inventory[i].GetItem().GetType()+
			" Price: ", slots.Inventory[i].GetItem().GetPrice(), " isAvailable: ", !slots.Inventory[i].IsSoldOut())
	}
}
