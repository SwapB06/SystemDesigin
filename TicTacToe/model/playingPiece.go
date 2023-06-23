package model

type IPlayingPiece interface {
	NewPlayingPiece()
}
type PlayingPiece struct {
	PieceType PieceType
}
