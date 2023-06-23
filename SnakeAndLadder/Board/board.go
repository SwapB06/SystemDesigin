package Board

import (
	"math/rand"
	"snakeandladder/Cell"
	"snakeandladder/Jump"
)

var Cells = [][]Cell.Cell{}

//var Cells = make([][]Cell.Cell, 0 /* 10*/)

type Board struct {
	BoardSize       int
	NumberOfSnakes  int
	NumberOfLadders int
}

func initializeCells(boardSize int) {
	Cells = make([][]Cell.Cell, boardSize)
	for i := 0; i < boardSize; i++ {
		Cells[i] = make([]Cell.Cell, boardSize)
		for j := 0; j < boardSize; j++ {
			c := Cell.Cell{
				Jump: Jump.Jump{},
			}
			Cells[i][j] = c
		}
	}

}

func addSnakesLadders(cells *[][]Cell.Cell, boardSize, numberOfSnakes, numberOfLadders int) {

	for numberOfSnakes > 0 {
		max := len(*cells)*boardSize /*len(cells[0])*/ - 1
		min := 1
		snakeHead := rand.Intn((max - min) + min)
		snakeTail := rand.Intn((max - min) + min)
		if snakeTail >= snakeHead {
			continue
		}

		snakeObj := /*new(Jump.Jump)*/ Jump.Jump{}
		snakeObj.Start = snakeHead
		snakeObj.End = snakeTail

		cell := GetCell(snakeHead)
		cell.Jump = snakeObj

		numberOfSnakes--
	}

	for numberOfLadders > 0 {
		max := len(*cells)*boardSize /*len(cells[0])*/ - 1
		min := 1
		ladderStart := rand.Intn((max - min) + min)
		ladderEnd := rand.Intn((max - min) + min)
		if ladderStart >= ladderEnd {
			continue
		}

		ladderObj := Jump.Jump{}
		ladderObj.Start = ladderStart
		ladderObj.End = ladderEnd

		cell := GetCell(ladderStart)
		cell.Jump = ladderObj

		numberOfLadders--
	}
}

func GetCell(playerPosition int) *Cell.Cell {
	boardRow := playerPosition / len(Cells)
	boardColumn := (playerPosition % len(Cells))
	return &Cells[boardRow][boardColumn]
}

func InitializeBoard(board Board) {
	initializeCells(board.BoardSize)
	addSnakesLadders(&Cells, board.BoardSize, board.NumberOfSnakes, board.NumberOfLadders)
}
