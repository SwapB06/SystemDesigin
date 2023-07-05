package controller

type VendingMachine struct {
	VendingMachineState State
	Inventory           Inventory
	CoinList            []Coin
}

func NewVendingMachine() VendingMachine {
	return VendingMachine{
		VendingMachineState: IdleState{},
		Inventory:           NewInventory(10),
		CoinList:            make([]Coin, 0),
	}
}

func (vm VendingMachine) GetVendingMachineState() State {
	return vm.VendingMachineState
}

func (vm *VendingMachine) SetVendingMachineState(vendingMachineState State) {
	vm.VendingMachineState = vendingMachineState
}
func (vm VendingMachine) GetInventory() *Inventory {
	return &vm.Inventory
}
func (vm *VendingMachine) SetInventory(inventory Inventory) {
	vm.Inventory = inventory
}
func (vm VendingMachine) GetCoinList() []Coin {
	return vm.CoinList
}
func (vm *VendingMachine) SetCoinList(coinList []Coin) {
	vm.CoinList = coinList
}
