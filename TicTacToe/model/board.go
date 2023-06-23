package model

import "fmt"

type Board struct {
	Size  int
	Board [][]*PlayingPiece
}

func NewBoard(size int) *Board {
	b := Board{
		Size:  size,
		Board: make([][]*PlayingPiece, size),
	}
	for i := 0; i < size; i++ {
		b.Board[i] = make([]*PlayingPiece, size)
	}
	return &b
}

func AddPiece(row, column int, playingPiece PlayingPiece) bool {
	if G.GameBoard.Board[row][column] != nil { //GameBoard.Board[row][column]
		return false
	}
	G.GameBoard.Board[row][column] = &playingPiece //GameBoard.Board[row][column]
	return true
}

type Pair struct {
	i int
	j int
}

func GetFreeCells() []Pair {
	freeCells := make([]Pair, 0)

	for i := 0; i < G.GameBoard.Size; i++ {
		for j := 0; j < G.GameBoard.Size; j++ {
			if G.GameBoard.Board[i][j] == nil {
				rowColumn := Pair{
					i: i,
					j: j,
				}
				freeCells = append(freeCells, rowColumn)
			}
		}
	}
	return freeCells
}

func printBoard() {
	for i := 0; i < G.GameBoard.Size; i++ {
		for j := 0; j < G.GameBoard.Size; j++ {
			if G.GameBoard.Board[i][j] != nil {
				fmt.Print(G.GameBoard.Board[i][j].PieceType, "   ")
			} else {
				fmt.Print("    ")
			}
			fmt.Print(" | ")
		}
		fmt.Println()
	}
}
