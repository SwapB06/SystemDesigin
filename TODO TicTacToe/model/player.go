package model

type Player struct {
	Name         string
	PlayingPiece PlayingPiece
}

func NewPlayer(name string, playingPiece PlayingPiece) *Player {
	return &Player{
		Name:         name,
		PlayingPiece: playingPiece,
	}
}

func (p Player) GetPlayerName() string {
	return p.Name
}

func (p *Player) SetPlayerName(name string) {
	p.Name = name
}

func (p Player) GetPlayingPiece() PlayingPiece {
	return p.PlayingPiece
}

func (p *Player) SetPlayingPiece(pp PlayingPiece) {
	p.PlayingPiece = pp
}
