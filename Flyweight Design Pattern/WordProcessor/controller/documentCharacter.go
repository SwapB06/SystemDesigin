package controller

import (
	"fmt"
)

type DocumentCharacter struct {
	Character rune
	FontType  string
	Size      int
}

func NewDocumentCharacter(char rune, font string, size int) *DocumentCharacter {
	return &DocumentCharacter{
		Character: char,
		FontType:  font,
		Size:      size,
	}
}

// only getter methods
func (d DocumentCharacter) Display(row, column int) {
	fmt.Printf(">> %c %s %d", d.Character, d.FontType, d.Size)
}
