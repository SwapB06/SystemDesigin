package Player

type Player struct {
	Id              string
	CurrentPosition int
}

func NewPlayer(id string, currPos int) *Player {
	return &Player{
		Id:              id,
		CurrentPosition: currPos,
	}
}

//getters and setters
