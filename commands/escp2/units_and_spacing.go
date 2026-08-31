package escp2

import "github.com/mjwhodur/imperfekt/commands/escp2/commandcodes"

func SetDefaultUnit(nominator uint8) []byte {
	if nominator == 5 || nominator == 10 || nominator == 20 || nominator == 30 || nominator == 40 || nominator == 50 || nominator == 60 {
		return []byte{commandcodes.ESC, 40, 85, 1, 0, nominator}
	} else {
		return []byte{commandcodes.ESC, 40, 85, 1, 0, 10}
	}

}

func Select1_8_inchLineSpacing() []byte {
	return []byte{commandcodes.ESC, '0'}
}

func Select1_6_inchLineSpacing() []byte {
	return []byte{commandcodes.ESC, '2'}
}
