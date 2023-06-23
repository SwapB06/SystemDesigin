package model

type PlayingPieceX struct {
	PieceType PieceType
}

func (x PlayingPieceX) NewPlayingPiece() PlayingPiece {
	return PlayingPiece{
		PieceType: X,
	}
}
