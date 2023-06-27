package controller

type cacheMap map[string]IRobot

type RoboticFactory struct {
	RoboticObjectCache cacheMap
}

func NewRobotFactory() *RoboticFactory {
	// var m cache
	// m = make(cache, 0)
	return &RoboticFactory{
		RoboticObjectCache: make(cacheMap, 0),
	}
}
func CreateRobot(robotType string) IRobot {
	c := NewRobotFactory()
	if _, ok := c.RoboticObjectCache[robotType]; ok {
		return c.RoboticObjectCache[robotType]
	} else {
		if robotType == "HUMANOID" {
			humanoidSprite := NewSprites()
			humanoidObject := NewHumanoidRobot(robotType, humanoidSprite)
			c.RoboticObjectCache[robotType] = humanoidObject
			return humanoidObject
		} else if robotType == "ROBOTICDOG" {
			roboticDogSprite := NewSprites()
			roboticDogObject := NewRoboticDog(robotType, roboticDogSprite)
			c.RoboticObjectCache[robotType] = roboticDogObject
			return roboticDogObject
		}
	}
	return nil
}
