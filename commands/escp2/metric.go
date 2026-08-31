package escp2

import "github.com/mjwhodur/imperfekt/commands/escp2/helpers"

func SetLeftMarginMilimeters(margin uint32) []byte {
	marginInInches := float64(margin) / 254
	marginInCols := marginInInches * 360 / 10

	return SetLeftMarginInCols15CPI(uint32(marginInCols))
}

// Sets margin from THE LEFT!
func SetRightMarginMillimeters(margin uint32) []byte {
	helpers.MilimeterToUnit(margin, 10)
	marginInInches := float64(margin) / 254
	marginInCols := marginInInches * 3600 / 10

	return SetRightMarginInCols15CPI(uint32(marginInCols))
}

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
	panic("Not implemented")
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

// Use when margin is known, set the absolute print position, counted without the margin (from the top of the sheet)
func SetAbsoluteVerticalPrintPositionInMillimetersWithKnownMargin(fromTop uint32, topMargin uint32) []byte {
	definedUnit := 10

	verPos, _ := helpers.MilimeterToUnit(fromTop, uint32(definedUnit))
	topMarginPos, _ := helpers.MilimeterToUnit(topMargin, uint32(definedUnit))

	upper := byte((verPos - topMarginPos) / 256)
	lower := byte((verPos - topMarginPos) % 256)
	return SetAbsoluteVerticalPrintPosition(lower, upper)
}

func SetAbsoluteHorizontalPrintPositionInMillimetersWithKnownMargin(fromLeft uint32, leftMargin uint32) []byte {
	definedUnit := 10

	verPos, _ := helpers.MilimeterToUnit(fromLeft, uint32(definedUnit))
	leftMarginPos, _ := helpers.MilimeterToUnit(leftMargin, uint32(definedUnit))

	upper := byte((verPos - leftMarginPos) / 256)
	lower := byte(((verPos - leftMarginPos) % 256))
	return SetAbsoluteHorizontalPosition(lower, upper)
}
func SetTopMarginMillimeters(nominator uint8) []byte {
	return []byte{}
}
