package controller

import "fmt"

type RoboticDog struct {
	Type string
	Body Sprites
}

func NewRoboticDog(typ string, body Sprites) RoboticDog {
	return RoboticDog{
		Type: typ,
		Body: body,
	}
}

func (rd RoboticDog) GetType() string {
	return rd.Type
}

func (rd RoboticDog) GetBody() Sprites {
	return rd.Body
}

func (rd RoboticDog) Display(x, y int) {
	fmt.Println("RoboticDog: ", x, y)
}
