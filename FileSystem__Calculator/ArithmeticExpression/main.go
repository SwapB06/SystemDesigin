package main

import "fmt"
import "calci/controller"

func main() {
	two := controller.NewNumber(2)

	one := controller.NewNumber(1)
	seven := controller.NewNumber(7)

	addExpression := controller.NewExpression(one, seven, controller.ADD)

	parentExpression := controller.NewExpression(two, addExpression, controller.MULTIPLY)

	fmt.Println(parentExpression.Evaluate())

}
