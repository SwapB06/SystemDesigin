package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	id              string
	currentPosition int
}

func newPlayer(id string, currPos int) *Player {
	return &Player{
		id:              id,
		currentPosition: currPos,
	}
}

type Jump struct {
	start int
	end   int
}

var diceCount int

const min int = 1
const max int = 6

func Dice(diceC int) {
	diceCount = diceC
}

func rollDice() int {
	totalSum := 0
	diceUsed := 0

	//diceCount = 1
	for diceUsed < diceCount {
		totalSum += rand.Intn((max - min) + min)
		diceUsed++
	}
	//fmt.Println(">>>", totalSum)
	return totalSum
}

type Cell struct {
	jump Jump
}

var cells = make([][]Cell /*board.boardSize*/, 10)

type Board struct {
	boardSize       int
	numberOfSnakes  int
	numberOfLadders int
}

func initializeCells(boardSize int) {
	for i := 0; i < boardSize; i++ {
		cells[i] = make([]Cell, boardSize)
		for j := 0; j < boardSize; j++ {
			c := Cell{
				jump: Jump{},
			}
			cells[i][j] = c
		}
	}

}
func addSnakesLadders(cells *[][]Cell, numberOfSnakes int, numberOfLadders int) {

	for numberOfSnakes > 0 {
		max := len(*cells)*board.boardSize /*len(cells[0])*/ - 1
		min := 1
		snakeHead := rand.Intn((max - min) + min)
		snakeTail := rand.Intn((max - min) + min)
		if snakeTail >= snakeHead {
			continue
		}

		snakeObj := /*new(Jump.Jump)*/ Jump{}
		snakeObj.start = snakeHead
		snakeObj.end = snakeTail

		cell := getCell(snakeHead)
		cell.jump = snakeObj

		numberOfSnakes--
	}

	for numberOfLadders > 0 {
		max := len(*cells)*board.boardSize /*len(cells[0])*/ - 1
		min := 1
		ladderStart := rand.Intn((max - min) + min)
		ladderEnd := rand.Intn((max - min) + min)
		if ladderStart >= ladderEnd {
			continue
		}

		ladderObj := Jump{}
		ladderObj.start = ladderStart
		ladderObj.end = ladderEnd

		cell := getCell(ladderStart)
		cell.jump = ladderObj

		numberOfLadders--
	}
}

func getCell(playerPosition int) *Cell {
	boardRow := playerPosition / len(cells)
	boardColumn := (playerPosition % len(cells))
	return &cells[boardRow][boardColumn]
}

var board Board
var PlayerList = make([]*Player, 0)
var winner *Player
var winnerStatus bool = false

func InitializeGame() {
	board = Board{
		boardSize:       10,
		numberOfSnakes:  5,
		numberOfLadders: 4,
	}
	Dice(1)
	//winner = nil
	addPlayers()
}
func addPlayers() {
	player1 := newPlayer("p1", 0)

	player2 := newPlayer("p2", 0)
	PlayerList = append(PlayerList, player1, player2)
}
func InitializeBoard(board Board) {
	initializeCells(board.boardSize)
	addSnakesLadders(&cells, board.numberOfSnakes, board.numberOfLadders)
}
func startGame() {
	for winnerStatus == false {

		//check whose turn now
		//playerTurn := findPlayerTurn()
		playerTurn := PlayerList[0]
		PlayerList = PlayerList[1:]
		PlayerList = append(PlayerList, playerTurn)

		fmt.Println("player turn is:"+playerTurn.id+" current position is: ",
			playerTurn.currentPosition)

		//roll the dice
		diceNumbers := rollDice()

		//get the new position
		playerNewPosition := playerTurn.currentPosition + diceNumbers
		playerNewPosition = jumpCheck(playerNewPosition)
		playerTurn.currentPosition = playerNewPosition

		fmt.Println("player turn is:"+playerTurn.id+" new Position is: ", playerNewPosition)
		//check for winning condition
		if playerNewPosition >= len(cells)*len(cells[0])-1 {

			winner = playerTurn
			winnerStatus = true
		}

	}
	fmt.Println("WINNER IS:" + winner.id)
}

//TODO: Try to keep this function
// func findPlayerTurn() *Player {

// 	playerTurn := PlayerList[0]
// 	PlayerList = PlayerList[1:]
// 	PlayerList = append(PlayerList, playerTurn)
// 	fmt.Println(">>>PlayerLIst", PlayerList)
// 	return playerTurn
// }

func jumpCheck(playerNewPosition int) int {

	if playerNewPosition > len(cells)*len(cells[0])-1 {
		return playerNewPosition
	}

	cell := getCell(playerNewPosition)
	fmt.Println(">>>>>", cell.jump.start)
	if /*cell.jump != nil &&*/ cell.jump.start == playerNewPosition {
		jumpBy := ""
		if cell.jump.start < cell.jump.end {
			jumpBy = "ladder"
		} else {
			jumpBy = "snake"
		}
		fmt.Println("jump done by: " + jumpBy)
		return cell.jump.end
	}
	return playerNewPosition /* + 1 */
}
func main() {
	InitializeGame()
	InitializeBoard(board)
	startGame()
}

/*we can have model directory and all object files in it, that will optimize import because in that case only one
import statement will be needed*/
