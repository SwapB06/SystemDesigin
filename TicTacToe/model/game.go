package model

import (
	"fmt"
	"strconv"
	"strings"
)

type PlayerList []*Player
type Game struct {
	PL        PlayerList
	GameBoard *Board
}

var G = &Game{}

func InitializeGame() {
	p := make([]*Player, 0)
	G = &Game{
		PL: p,
	}
	x := PlayingPieceX{}
	o := PlayingPieceO{}
	crossPiece := x.NewPlayingPiece()
	noughtPiece := o.NewPlayingPiece()

	player1 := NewPlayer("player1", crossPiece)
	player2 := NewPlayer("Player2", noughtPiece)

	G.PL = append(G.PL, player1)
	G.PL = append(G.PL, player2)

	G.GameBoard = NewBoard(3)
}

func StartGame() string {
	noWinner := true
	for noWinner {
		//take out the player whose turn is and also put the player in the list back
		playerTurn := G.PL[0]
		G.PL = G.PL[1:]

		//get the free space from the board
		printBoard()
		freeSpaces := GetFreeCells()
		if len(freeSpaces) == 0 {
			noWinner = false
			continue
		}

		//read the user input
		fmt.Println("Player:", playerTurn.Name, " Enter row,column: ")
		var s string
		fmt.Scanln(&s)
		values := strings.Split(s, ",")
		inputRow, _ := strconv.Atoi(values[0])
		inputColumn, _ := strconv.Atoi(values[1])

		//place the piece
		pieceAddedSuccessfully := AddPiece(inputRow, inputColumn, playerTurn.PlayingPiece)
		if !pieceAddedSuccessfully {
			//player can not insert the piece into this cell, player has to choose another cell
			fmt.Println("Incorredt possition chosen, try again")
			G.PL = append([]*Player{playerTurn}, G.PL...)
			//G.PL.addFirst(playerTurn)
			continue
		}
		G.PL = append(G.PL, playerTurn)

		winner := isThereWinner(inputRow, inputColumn, playerTurn.PlayingPiece.PieceType)
		if winner {
			return playerTurn.Name
		}
	}
	return "tie"
}

func isThereWinner(row, column int, pieceType PieceType) bool {

	rowMatch := true
	columnMatch := true
	diagonalMatch := true
	antiDiagonalMatch := true

	//need to check in row
	for i := 0; i < G.GameBoard.Size; i++ {

		if G.GameBoard.Board[row][i] == nil || G.GameBoard.Board[row][i].PieceType != pieceType {
			rowMatch = false
		}
	}

	//need to check in column
	for i := 0; i < G.GameBoard.Size; i++ {

		if G.GameBoard.Board[i][column] == nil || G.GameBoard.Board[i][column].PieceType != pieceType {
			columnMatch = false
		}
	}

	//need to check diagonals
	for i, j := 0, 0; i < G.GameBoard.Size; i, j = i+1, j+1 {
		if G.GameBoard.Board[i][j] == nil || G.GameBoard.Board[i][j].PieceType != pieceType {
			diagonalMatch = false
		}
	}

	//need to check anti-diagonals
	for i, j := 0, G.GameBoard.Size-1; i < G.GameBoard.Size; i, j = i+1, j-1 {
		if G.GameBoard.Board[i][j] == nil || G.GameBoard.Board[i][j].PieceType != pieceType {
			antiDiagonalMatch = false
		}
	}

	return rowMatch || columnMatch || diagonalMatch || antiDiagonalMatch
}
