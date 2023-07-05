package controller

type State interface {
	ClickOnInsertCoinButton(machine *VendingMachine)
	ClickOnStartProductSelectionButton(machine *VendingMachine)
	InsertCoin(machine *VendingMachine, coin CoinType)
	ChooseProduct(machine *VendingMachine, codeNumber int)
	GetChange(returnChangeMoney int) int
	DispenseProduct(machine *VendingMachine, codeNumber int) Item
	RefundFullMoney(machine *VendingMachine) []Coin
	UpdateInventory(machine *VendingMachine, item Item, codeNumber int)
}
