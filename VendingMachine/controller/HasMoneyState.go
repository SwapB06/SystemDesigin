package controller

import (
	"fmt"
)

type HasMoneyState struct{}

func NewHasMoneyState() HasMoneyState {
	return HasMoneyState{}
}
func HasMoneyStateFun() {
	fmt.Println("Currently Vending machine is in HasMoneyState")
}
func (h HasMoneyState) ClickOnInsertCoinButton(machine *VendingMachine) {
	return
}
func (h HasMoneyState) ClickOnStartProductSelectionButton(machine *VendingMachine) {
	machine.SetVendingMachineState(SelectionState{})
}
func (h HasMoneyState) InsertCoin(machine *VendingMachine, coin CoinType) {
	fmt.Println("Accepted the coin")
	machine.CoinList = append(machine.GetCoinList(), Coin{Ct: coin, Value: int(coin)})
}
func (h HasMoneyState) ChooseProduct(machine *VendingMachine, codeNumber int) {
	fmt.Println("you need to click on start product selection button first")
}
func (h HasMoneyState) GetChange(returnChangeMoney int) int {
	fmt.Println("you can not get change in hasMoney state")
	return 0
}
func (h HasMoneyState) DispenseProduct(machine *VendingMachine, codeNumber int) Item {
	fmt.Println("product can not be dispensed in hasMoney state")
	return Item{}
}
func (h HasMoneyState) RefundFullMoney(machine *VendingMachine) []Coin {
	fmt.Println("Returned the full amount back in the Coin Dispense Tray")
	machine.SetVendingMachineState(IdleState{})
	return machine.GetCoinList()
}
func (h HasMoneyState) UpdateInventory(machine *VendingMachine, item Item, codeNumber int) {
	fmt.Println("You can not update inventory in hasMoney  state")
}
