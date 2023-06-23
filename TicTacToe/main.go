package main

import (
	"fmt"
	"tictactoe/model"
)

func main() {
	//model.NewTicTacGame()
	model.InitializeGame()
	fmt.Println("Game Winner is: ", model.StartGame())
}
