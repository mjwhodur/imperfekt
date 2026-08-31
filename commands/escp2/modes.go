package escp2

import "github.com/mjwhodur/imperfekt/commands/escp2/commandcodes"

func InitializePrinter() []byte {
	return []byte{commandcodes.ESC, 64}
}
