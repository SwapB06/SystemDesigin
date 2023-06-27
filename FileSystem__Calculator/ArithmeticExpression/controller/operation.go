package controller

type Operation int

// const (
// 	ADD      Operation = "+"
// 	SUBTRACT Operation = "-"
// 	MULTIPLY Operation = "*"
// 	DIVIDE   Operation = "/"
// )

const (
	ADD Operation = iota + 1
	SUBTRACT
	MULTIPLY
	DIVIDE
)
