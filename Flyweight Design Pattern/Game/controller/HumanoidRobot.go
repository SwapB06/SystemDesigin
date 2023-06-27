package controller

import "fmt"

type HumanoidRobot struct {
	Type string
	Body Sprites
}

func NewHumanoidRobot(typ string, body Sprites) HumanoidRobot {
	return HumanoidRobot{
		Type: typ,
		Body: body,
	}
}

func (hr HumanoidRobot) GetType() string {
	return hr.Type
}

func (hr HumanoidRobot) GetBody() Sprites {
	return hr.Body
}

func (hr HumanoidRobot) Display(x, y int) {
	fmt.Println("HumanoidRobot: ", x, y)
}
