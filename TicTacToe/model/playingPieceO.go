package model

type PlayingPieceO struct {
	PieceType PieceType
}

func (o PlayingPieceO) NewPlayingPiece() PlayingPiece {
	return PlayingPiece{
		PieceType: O,
	}
}
