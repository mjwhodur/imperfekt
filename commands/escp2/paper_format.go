package escp2

import "github.com/mjwhodur/imperfekt/commands/escp2/commandcodes"

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
func SetPageFormat(topLower byte, topHigher byte, bottomLower byte, bottomHigher byte) []byte {

	return []byte{commandcodes.ESC, 40, 99, 4, 0, topLower, topHigher, bottomLower, bottomHigher}
}

func SetPageLengthInLines(amount byte) []byte {
	return []byte{commandcodes.ESC, 67, amount}
}

func SetPageLengthInInches(length uint8) []byte {
	if length >= 1 && length <= 22 {
		return []byte{
			commandcodes.ESC, 67, commandcodes.NUL, length,
		}
	}
	panic("Page length must be between 1 and 22 inches")
}

func SetTopMargin(nominator uint8) []byte {
	panic("SetTopMargin not implemented")
	return []byte{
		commandcodes.ESC, 67, commandcodes.NUL, nominator,
	}
}

// SetBottomMargin sets the bottom margin on continuous paper to n lines in the current-set line spacing from the
// top of form position on the next page
func SetBottomMargin(nominator uint8) []byte {
	return []byte{
		commandcodes.ESC, 67, 'N', nominator,
	}
}

func CancelBottomMargin() []byte { return []byte{commandcodes.ESC, 79} }

// FIXME: Test
func SetLeftMarginInCols15CPI(marginInCols uint32) []byte {
	return []byte{commandcodes.ESC, 103, commandcodes.ESC, 'I', byte(marginInCols)}
}

// FIXME: Test
func SetRightMarginInCols15CPI(marginInCols uint32) []byte {
	return []byte{commandcodes.ESC, 103, commandcodes.ESC, 'Q', byte(marginInCols)}
}

// FIXME: Test
func SetLeftMarginInCols(marginInCols uint32) []byte {
	return []byte{commandcodes.ESC, 'I', byte(marginInCols)}
}

// FIXME: Test
func SetRightMarginInCols(marginInCols uint32) []byte {
	return []byte{commandcodes.ESC, 'Q', byte(marginInCols)}
}

func SetHorizontalTabs(tabs []byte) []byte {
	return append(append([]byte{commandcodes.ESC, 68}, tabs...), 0)
}

func SetVerticalTabs(tabs []byte) []byte {
	return append(append([]byte{commandcodes.ESC, 66}, tabs...), 0)
}

func SetN360thsinchLineSpacing(amount byte) []byte {
	return []byte{commandcodes.ESC, 43, amount}
}
