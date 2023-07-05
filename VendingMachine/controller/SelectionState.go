package controller

import "fmt"

type SelectionState struct{}

func NewSelectionState() SelectionState {
	return SelectionState{}
}
func SelectionStateFun() {
	fmt.Println("Currently Vending machine is in SelectionState")
}
func (s SelectionState) ClickOnInsertCoinButton(machine *VendingMachine) {
	fmt.Println("you can not click on insert coin button in Selection state")
}
func (s SelectionState) ClickOnStartProductSelectionButton(machine *VendingMachine) {
	return
}
func (s SelectionState) InsertCoin(machine *VendingMachine, coin CoinType) {
	fmt.Println("you can not insert Coin in selection state")
}
func (s SelectionState) ChooseProduct(machine *VendingMachine, codeNumber int) {
	//1. get item of this codeNumber
	item := machine.GetInventory().GetItem(codeNumber)
	paidByUser := 0

	//2. total amount paid by User
	for _, coin := range machine.GetCoinList() {
		paidByUser = paidByUser + coin.Value
	}

	//3. compare product price and amount paid by user
	if paidByUser < item.GetPrice() {
		fmt.Println("Insufficient Amount, Product you selected is for price: ", item.GetPrice(), " and you paid: ", paidByUser)
		_ = s.RefundFullMoney(machine)
		fmt.Println("Insufficient amount")
	} else if paidByUser >= item.GetPrice() {
		if paidByUser > item.GetPrice() {
			_ = s.GetChange(paidByUser - item.GetPrice())
		}
		ds := NewDespenseState()
		machine.SetVendingMachineState(ds)
	}
}
func (s SelectionState) GetChange(returnExtraMoney int) int {
	//actual logic should be to return COINs in the dispense tray, but for simplicity i am just returning the amount to be refunded
	fmt.Println("Returned the change in the Coin Dispense Tray: ", returnExtraMoney)
	return returnExtraMoney
}
func (s SelectionState) DispenseProduct(machine *VendingMachine, codeNumber int) Item {
	fmt.Println("product can not be dispensed Selection state")
	return Item{}
}
func (s SelectionState) RefundFullMoney(machine *VendingMachine) []Coin {
	fmt.Println("Returned the full amount back in the Coin Dispense Tray")
	machine.SetVendingMachineState(NewIdleState())
	return machine.GetCoinList()
}
func (s SelectionState) UpdateInventory(machine *VendingMachine, item Item, codeNumber int) {
	fmt.Println("Inventory can not be updated in Selection state")
}
