package main

import "wordprocessor/controller"

func main() {
	obj1 := controller.CreateLetter('t')
	obj1.Display(0, 0)

	obj2 := controller.CreateLetter('t')
	obj2.Display(0, 6)
}
