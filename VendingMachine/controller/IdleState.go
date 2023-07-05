package controller

import (
	"fmt"
)

type IdleState struct {
}

func NewIdleState() IdleState {
	return IdleState{}
}
func IdleStateFun() {
	fmt.Println("Currently Vending machine is in IdleState")
}

//	func IdleState(vm controller.VendingMachine) {
//		fmt.Println("Currently Vending machine is in IdleState")
//		vm.SetCoinList(make([]controller.Coin, 0))
//	}
func (i IdleState) ClickOnInsertCoinButton(vm *VendingMachine) {
	vm.SetVendingMachineState(NewHasMoneyState())
}

func (i IdleState) ClickOnStartProductSelectionButton(vm *VendingMachine) {
	fmt.Println("first you need to click on insert coin button")

}

func (i IdleState) InsertCoin(vm *VendingMachine, coin CoinType) {
	fmt.Println("You can not insert coin in idle state")
}

func (i IdleState) ChooseProduct(vm *VendingMachine, codeNumber int) {
	fmt.Println("You cannot choose product in idle state")
}

func (i IdleState) GetChange(returnChangeMoney int) int {
	fmt.Println("You cannpt get change in idle state")
	return 0
}

func (i IdleState) RefundFullMoney(vm *VendingMachine) []Coin {
	fmt.Println("You cannot get refunded in idle state")
	return make([]Coin, 0)
}

func (i IdleState) DispenseProduct(vm *VendingMachine, codeNumber int) Item {
	fmt.Println("Product cannot be dispensed in idle state")
	return Item{}
}

func (i IdleState) UpdateInventory(vm *VendingMachine, item Item, codeNumber int) {
	vm.GetInventory().AddItem(item, codeNumber)
}
