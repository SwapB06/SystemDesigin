package controller

type cache map[rune]ILetter

type LetterFactor struct {
	CharCache cache
}

func NewLetterFactor() *LetterFactor {
	// var m cache
	// m = make(cache, 0)
	return &LetterFactor{
		CharCache: make(cache, 0),
	}
}
func CreateLetter(characterValue rune) ILetter {
	c := NewLetterFactor()
	if _, ok := c.CharCache[characterValue]; ok {
		return c.CharCache[characterValue]
	} else {
		charObj := NewDocumentCharacter(characterValue, "Arial", 10)
		c.CharCache[characterValue] = charObj
		return charObj
	}
}
