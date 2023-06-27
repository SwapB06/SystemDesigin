package controller

import "fmt"

type Number struct {
	Value int
}

func NewNumber(value int) *Number {
	return &Number{
		Value: value,
	}
}
func (n Number) Evaluate() int {
	fmt.Println("Number value is :", n.Value)
	return n.Value
}
