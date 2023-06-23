package Dice

import "math/rand"

var diceCount int

const min int = 1
const max int = 6

func Dice(diceC int) {
	diceCount = diceC
}

func RollDice() int {
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
