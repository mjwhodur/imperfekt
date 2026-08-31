package escp2

import (
	"github.com/mjwhodur/imperfekt/commands/escp2/commandcodes"
	"github.com/mjwhodur/imperfekt/commands/escp2/fonts"
)

func SetQualityLQ() []byte {
	return []byte{commandcodes.ESC, 120, 1}
}

func SetQualityDraft() []byte {
	return []byte{commandcodes.ESC, 120, 0}
}

func SelectFont(font fonts.Font) []byte {
	return []byte{commandcodes.ESC, 107, byte(font)}
}

func Select10andhalfPoint10CPI() []byte { return []byte{commandcodes.ESC, 80} }

func Select10andhalfPoint12CPI() []byte { return []byte{commandcodes.ESC, 77} }

func Select10andhalfPoint15CPI() []byte { return []byte{commandcodes.ESC, 103} }

func SelectBoldFont() []byte   { return []byte{commandcodes.ESC, 69} }
func CancelBoldFont() []byte   { return []byte{commandcodes.ESC, 70} }
func SelectItalicFont() []byte { return []byte{commandcodes.ESC, 52} }
func CancelItalicFont() []byte { return []byte{commandcodes.ESC, 53} }

func SelectDoubleStrikePrinting() []byte { return []byte{commandcodes.ESC, 71} }
func CancelDoubleStrikePrinting() []byte { return []byte{commandcodes.ESC, 72} }

func EnableUnderline() []byte  { return []byte{commandcodes.ESC, 45, 1} }
func DisableUnderline() []byte { return []byte{commandcodes.ESC, 45, 0} }

func SelectSuperScript() []byte    { return []byte{commandcodes.ESC, 83, 0} }
func SelectSubScript() []byte      { return []byte{commandcodes.ESC, 83, 1} }
func CancelSuperSubScript() []byte { return []byte{commandcodes.ESC, 84} }

func MasterSelect(value byte) []byte {
	return []byte{commandcodes.ESC, 33, value}
}

func TurnUnderlineOnOff(value byte) []byte {
	return []byte{commandcodes.ESC, 45, value}
}

func SelectLineScore(par1 byte, par2 byte) []byte {
	return []byte{commandcodes.ESC, 40, 45, 3, 0, 1, par1, par2}
}

func SelectCharacterStyle(param byte) []byte {
	return []byte{commandcodes.ESC, param}
}

func SelectCondensedPrinting() []byte {
	return []byte{commandcodes.SI}
}

func SelectCondensedPrinting2() []byte {
	return []byte{commandcodes.ESC, commandcodes.SI}
}

func CancelCondensedPrinting() []byte {
	return []byte{commandcodes.DC2}
}

func SelectDoubleWidthPrintingOneLine() []byte {
	return []byte{14}
}

func SelectDoubleWidthPrintingOneLine2() []byte {
	return []byte{commandcodes.ESC, 14}
}

func CancelDoubleWidthPrintingOneLine() []byte {
	return []byte{commandcodes.DC4}
}

func TurnProportionalModeOnOff(value byte) []byte {
	return []byte{commandcodes.ESC, 112, value}
}

func TurnDoubleWidthPrintingOnOff(param byte) []byte {
	return []byte{commandcodes.ESC, 87, param}
}

func TurnDoubleHeightPrintingOnOff(param byte) []byte {
	return []byte{commandcodes.ESC, 119, param}
}

func SelectFontByPitchAndPoint(pitch byte, pointLower byte, pointHigher byte) []byte {
	return []byte{
		commandcodes.ESC, pitch, pointLower, pointHigher,
	}
}

// SetIntercharacterSpace increases the space between characters by value/180 inch in LQ mode and n/120 inch in draft mode
func SetIntercharacterSpace(value byte) []byte {
	return []byte{commandcodes.ESC, 32, value}
}
