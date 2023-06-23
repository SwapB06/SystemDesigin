package main

import "snakeandladder/Game"

func main() {
	Game.InitializeGame()
	//Game.InitializeBoard(board)
	Game.StartGame()
}

/*we can have model directory and all object files in it, that will optimize import because in that case only one
import statement will be needed*/
