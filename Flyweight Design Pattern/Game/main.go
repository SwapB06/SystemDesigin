package main

import "game/controller"

func main() {
	humanoidRobot1 := controller.CreateRobot("HUMANOID")
	humanoidRobot1.Display(1, 2)

	humanoidRobot2 := controller.CreateRobot("HUMANOID")
	humanoidRobot2.Display(10, 30)

	roboDog1 := controller.CreateRobot("ROBOTICDOG")
	roboDog1.Display(2, 9)

	roboDog2 := controller.CreateRobot("ROBOTICDOG")
	roboDog2.Display(11, 19)
}
