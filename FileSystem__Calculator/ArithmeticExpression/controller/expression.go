package controller

import "fmt"

type Expression struct {
	LeftExpression  ArithmeticExpression
	RightExpression ArithmeticExpression
	Op              Operation
}

func NewExpression(leftPart, rightPart ArithmeticExpression, op Operation) *Expression {
	return &Expression{
		LeftExpression:  leftPart,
		RightExpression: rightPart,
		Op:              op,
	}
}

func (e Expression) Evaluate() int {
	value := 0
	switch e.Op {
	case ADD:
		value = e.LeftExpression.Evaluate() + e.RightExpression.Evaluate()
	case SUBTRACT:
		value = e.LeftExpression.Evaluate() - e.RightExpression.Evaluate()
	case MULTIPLY:
		value = e.LeftExpression.Evaluate() * e.RightExpression.Evaluate()
	case DIVIDE:
		value = e.LeftExpression.Evaluate() / e.RightExpression.Evaluate()
	}
	fmt.Println("Expression value is :", value)
	return value
}
