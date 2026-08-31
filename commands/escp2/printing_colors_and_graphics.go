package escp2

import "github.com/mjwhodur/imperfekt/commands/escp2/commandcodes"

func SelectGraphicsMode()        {}
func SelectMicroWeavePrintMode() {}
func PrintRasterGraphics()       {}
func EnterTIFFCompressedMode()   {}

func SelectBitImage(density byte, lower byte, upper byte, graphicsData []byte) []byte {
	return append([]byte{commandcodes.ESC, 42, density, lower, upper}, graphicsData...)
}

func ReassignBitImageMode()            {}
func Select60DPIGraphics()             {}
func Select120DPIGraphics()            {}
func Select120DPIDoubleSpeedGraphics() {}
func Select240DPIGraphics()            {}

func SelectPrintingColor(color byte) []byte {
	return []byte{commandcodes.ESC, 114, color}
}
