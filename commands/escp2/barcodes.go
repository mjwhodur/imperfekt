package escp2

import "github.com/mjwhodur/imperfekt/commands/escp2/commandcodes"

func BarCodeSetupAndPrint(lower byte, upper byte, codeType byte, moduleWidth byte, spaceAdjustmentValue byte, barLengthv1 byte, barLenghtv2 byte, controlFlag byte, barcodeData []byte) []byte {
	return append([]byte{commandcodes.ESC, 40, 66, lower, upper, codeType, moduleWidth, spaceAdjustmentValue, barLengthv1, barLenghtv2, controlFlag}, barcodeData...)
}
