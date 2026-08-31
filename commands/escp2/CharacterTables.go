package escp2

import (
	"github.com/mjwhodur/imperfekt/commands/escp2/charmaps"
)

func SelectCharacterTable(which byte) []byte {
	if (which <= 3) || (48 <= which && which <= 51) {
		return []byte{27, 116, which}
	}
	return []byte{27, 116, 0}
}

// AssignCharacterTable assigns the character table to the selected memory table
func AssignCharacterTable(which byte, characterTable charmaps.Chartable) []byte {
	header := []byte{27, 40, 116, 3, 0}
	if (which <= 3) || (48 <= which && which <= 51) {
		return append(append(header, which), charmaps.GetTableCode(characterTable)...)
	}
	return nil
}
