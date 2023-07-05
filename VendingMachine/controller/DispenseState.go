package controller

import "fmt"

type DispenseState struct {
}

func NewDespenseState() DispenseState {
	return DispenseState{}
}
func DispenseStateFun(vm *VendingMachine, codeNumber int) {
	fmt.Println("Currently vending machine is in DispenseState")
	NewDespenseState().DispenseProduct(vm, codeNumber)
}

func (d DispenseState) ClickOnInsertCoinButton(vm *VendingMachine) {
	fmt.Println("insert coin button can not clicked on Dispense state")
}

func (d DispenseState) ClickOnStartProductSelectionButton(vm *VendingMachine) {
	fmt.Println("product selection buttion can not be clicked in Dispense state")
}

func (d DispenseState) InsertCoin(vm *VendingMachine, coin CoinType) {
	fmt.Println("coin can not be inserted in Dispense state")
}

func (d DispenseState) ChooseProduct(vm *VendingMachine, codeNumber int) {
	fmt.Println("product can not be choosen in Dispense state")
}

func (d DispenseState) GetChange(returnChangeMoney int) int {
	fmt.Println("change can not returned in Dispense state")
	return 0
}

func (d DispenseState) RefundFullMoney(vm *VendingMachine) []Coin {
	fmt.Println("refund can not be happen in Dispense state")
	return make([]Coin, 0)
}

func (d DispenseState) DispenseProduct(vm *VendingMachine, codeNumber int) Item {
	fmt.Println("Product has been dispensed")
	item := vm.GetInventory().GetItem(codeNumber)
	vm.GetInventory().UpdateSoldOutItem(codeNumber)
	vm.SetVendingMachineState(NewIdleState())
	return item
}

func (d DispenseState) UpdateInventory(vm *VendingMachine, item Item, codeNumber int) {
	fmt.Println("inventory can not be updated in Dispense state")
}
