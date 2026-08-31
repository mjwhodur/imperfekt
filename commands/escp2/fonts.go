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
