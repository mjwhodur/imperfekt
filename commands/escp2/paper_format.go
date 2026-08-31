package escp2

import "github.com/mjwhodur/imperfekt/commands/escp2/helpers"

func SetPageSizeInMillimeters(length uint32, width uint32) []byte {
	// 1 inch = 254 mm

	// Assuming default unit is set to 10 (1 unit = 1 / 360")
	setDefaultUnit := SetDefaultUnit(10)
	// Page Length in Inches = paper_in_mm / 254

	// In units: 1 inch = 360 units

	// Epson_LTH_in_units = paper_length_in_inches * 360

	EpsonLthInUnits, err := helpers.MilimeterToUnit(length, 10)
	if err != nil {
		panic(err)
	}
	pageLen := SetPageLengthInDefinedUnit(EpsonLthInUnits)
	//pageWidth := SetRightMarginMillimeters(width)
	return append(append(setDefaultUnit, pageLen...))

}

func SetPageMarginsInMilimeters(top uint32, bottom uint32, left uint32, right uint32) []byte {
	return nil
}

func SetTopBottomMarginsInMilimeters(top uint32, bottom uint32) []byte {
	definedUnit := 10
	topMargin, _ := helpers.MilimeterToUnit(top, uint32(definedUnit))
	bottomMargin, _ := helpers.MilimeterToUnit(bottom, uint32(definedUnit))

	return SetPageFormat(topMargin, bottomMargin, byte(definedUnit))
}

func SetPageFormatA4() []byte {
	return SetPageSizeInMillimeters(297, 210)
}

func SetPageFormatC6() []byte {
	return SetPageSizeInMillimeters(162, 114)
}
