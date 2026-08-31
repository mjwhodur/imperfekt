package escp2

import (
	"github.com/mjwhodur/imperfekt/commands/escp2/commandcodes"
	"github.com/mjwhodur/imperfekt/commands/escp2/helpers"
)

// SetDefaultUnit sets measuring unit to x / 3600 inch, however different settings may alter the currently used
// unit in the printer.
// For example:
//
//			using the move to horizontally / vertically sets this value to 10
//	     	using letter quality (LQ) : 20 / 3600" (1/180")
//
// When passing improper value this command resets to standard 1/360"

func SetDefaultUnit(nominator uint8) []byte {
	if nominator == 5 || nominator == 10 || nominator == 20 || nominator == 30 || nominator == 40 || nominator == 50 || nominator == 60 {
		return []byte{commandcodes.ESC, 40, 85, 1, 0, nominator}
	} else {
		return []byte{commandcodes.ESC, 40, 85, 1, 0, 10}
	}

}

func SetPageLengthInDefinedUnit(length uint32, definedUnit ...uint8) []byte {
	//unit := uint8(10)
	//if len(definedUnit) > 0 {
	//	unit = definedUnit[0]
	//}

	valH := byte((length) / 256)
	valL := byte((length) % 256)

	return []byte{commandcodes.ESC, 40, 67, 2, 0, valL, valH}
}

// SetPageFormat sets the top and bottom margins in the defined units
func SetPageFormat(topMargin uint32, bottomMargin uint32, definedUnit byte) []byte {
	panic("Does not work in LQ-350")
	topLower := byte(uint32(float32(topMargin/uint32(definedUnit))) / 256)
	topHigher := byte(uint32(float32(topMargin/uint32(definedUnit))) % 256)
	bottomLower := byte(uint32(float32(bottomMargin/uint32(definedUnit))) / 256)
	bottomHigher := byte(uint32(float32(bottomMargin/uint32(definedUnit))) % 256)

	return []byte{commandcodes.ESC, 40, 99, 4, 0, topLower, topHigher, bottomLower, bottomHigher}
}

func SetPageLengthInLines(nominator uint8) []byte { return nil }

func SetPageLengthInInches(length uint8) []byte {
	if length >= 1 && length <= 22 {
		return []byte{
			commandcodes.ESC, 67, commandcodes.NUL, length,
		}
	}
	panic("Page length must be between 1 and 22 inches")
}

func SetTopMargin(nominator uint8) []byte {
	return nil
}

func SetTopMarginMillimeters(nominator uint8) []byte {
	return []byte{}
}

func SetBottomMargin(nominator uint8) []byte { return nil }

func CancelBottomMargin(nominator uint8) []byte { return nil }

func SetLeftMarginInCols(marginInCols uint32) []byte {
	// 1 col = 1/15 inch
	return []byte{commandcodes.ESC, 103, commandcodes.ESC, 'I', byte(marginInCols)}
}
func SetRightMarginInCols(marginInCols uint32) []byte {
	// 1 col = 1/15 inch = 24 / 360 inch
	return []byte{commandcodes.ESC, 103, commandcodes.ESC, 'Q', byte(marginInCols)}
}

func SetLeftMarginMilimeters(margin uint32) []byte {
	marginInInches := float64(margin) / 254
	marginInCols := marginInInches * 360 / 10

	return SetLeftMarginInCols(uint32(marginInCols))
}

// Sets margin from THE LEFT!
func SetRightMarginMillimeters(margin uint32) []byte {
	helpers.MilimeterToUnit(margin, 10)
	marginInInches := float64(margin) / 254
	marginInCols := marginInInches * 3600 / 10

	return SetRightMarginInCols(uint32(marginInCols))
}
