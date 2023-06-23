package Game

import (
	"fmt"
	"snakeandladder/Board"
	"snakeandladder/Dice"
	"snakeandladder/Player"
)

var board Board.Board

// var dice Dice.Dice
var PlayerList = make([]*Player.Player, 0)
var winner *Player.Player
var winnerStatus bool

func InitializeGame() {
	board = Board.Board{
		BoardSize:       10,
		NumberOfSnakes:  6,
		NumberOfLadders: 4,
	}
	Dice.Dice(1)
	Board.InitializeBoard(board)
	winnerStatus = false
	addPlayers()
}
func addPlayers() {
	player1 := Player.NewPlayer("p1", 0)

	player2 := Player.NewPlayer("p2", 0)
	player3 := Player.NewPlayer("p3", 0)
	PlayerList = append(PlayerList, player1, player2, player3)
}
func StartGame() {
	for winnerStatus == false {

		//check whose turn now
		//playerTurn := findPlayerTurn()
		playerTurn := PlayerList[0]
		PlayerList = PlayerList[1:]
		PlayerList = append(PlayerList, playerTurn)

		fmt.Println("player turn is:"+playerTurn.Id+" current position is: ",
			playerTurn.CurrentPosition)

		//roll the dice
		diceNumbers := Dice.RollDice()

		//get the new position
		playerNewPosition := playerTurn.CurrentPosition + diceNumbers
		playerNewPosition = jumpCheck(playerNewPosition)
		playerTurn.CurrentPosition = playerNewPosition

		fmt.Println("player turn is:"+playerTurn.Id+" new Position is: ", playerNewPosition)
		//check for winning condition
		if playerNewPosition >= len(Board.Cells)*len(Board.Cells[0])-1 {

			winner = playerTurn
			winnerStatus = true
		}

	}
	fmt.Println("WINNER IS:" + winner.Id)
}

//FindPlayerTurn function

func jumpCheck(playerNewPosition int) int {

	if playerNewPosition > len(Board.Cells)*len(Board.Cells[0])-1 {
		return playerNewPosition
	}

	cell := Board.GetCell(playerNewPosition)
	fmt.Println(">>>>>", cell.Jump.Start)
	if /*cell.jump != nil &&*/ cell.Jump.Start == playerNewPosition {
		jumpBy := ""
		if cell.Jump.Start < cell.Jump.End {
			jumpBy = "ladder"
		} else {
			jumpBy = "snake"
		}
		fmt.Println("jump done by: " + jumpBy)
		return cell.Jump.End
	}
	return playerNewPosition /* + 1 */
}
