package escp2

import (
	"github.com/mjwhodur/imperfekt/commands/escp2/commandcodes"
)

// The resulting horiszontal position is determined by the formula below:
// upper := ((horizontal - leftMargin) / definedUnit) // 256
// lower := ((horizontal - leftMargin) / definedUnit) % 256
//

func SetAbsoluteHorizontalPosition(lower byte, upper byte) []byte {
	return []byte{commandcodes.ESC, 36, lower, upper}
}

// SetAbsoluteVerticalPosition moves the vertical print position specified by the formula
func SetAbsoluteVerticalPrintPosition(lower byte, upper byte) []byte {
	return []byte{commandcodes.ESC, 40, 86, 2, 0, lower, upper}
}

func SetRelativeHorizontalPosition(lower byte, upper byte) []byte {
	return []byte{commandcodes.ESC, 92, lower, upper}
}

func SetRelativeVerticalPosition(lower byte, upper byte) []byte {
	return []byte{commandcodes.ESC, 40, 118, 2, 0, lower, upper}
}

func AdvancePrintPositionVertically(amount byte) []byte {
	return []byte{commandcodes.ESC, 74, amount}
}

func TabHorizontally() []byte {
	return []byte{commandcodes.HT}
}

func TabVertically() []byte {
	return []byte{commandcodes.VT}
}

// CarriageReturn moves the print position to the left-margin position.
// Always send a CR Command at the end of each line of text or graphics data.
func CarriageReturn() []byte {
	return []byte{commandcodes.CR}
}

// LineFeed advances the vertical print position one line (in the currently set line spacing).
// Moves the horizontal print position to the left-margin position.
// Ejects the paper if the print position ends below the bottom-margin position.
func LineFeed() []byte {
	return []byte{commandcodes.LF}
}

// FormFeed advances the vertical print position on continuous paper to the top-margin position of the next page.
// Ejects single-sheet paper. Moves the horizontal print position to the left-margin position. Prints all data in the buffer.
// It is recommended to always send a CR command before the FF command.
func FormFeed() []byte {
	return []byte{commandcodes.FF}
}
