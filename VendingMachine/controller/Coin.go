package controller

type CoinType int

const (
	PENNY   CoinType = 1
	NICKEL           = 5
	DIME             = 10
	QUARTER          = 25
)

type Coin struct {
	Ct    CoinType
	Value int
}

func NewCoin(value int) *Coin {
	return &Coin{
		Value: value,
	}
}
