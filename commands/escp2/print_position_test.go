package escp2

import "testing"

func TestSetAbsoluteHorizontalPosition1inchFromTheLeft(t *testing.T) {

	milimeters := 25.4
	margin := 0

	output := SetAbsoluteHorizontalPrintPositionInMillimetersWithKnownMargin(uint32(milimeters), uint32(margin))

	if output[2] != 104 && output[3] != 1 {
		t.Errorf("Expected: 1, 104, Got: %d, %d", output[3], output[2])
	}

}

func TestSetAbsoluteHorizontalPosition1inchFromTheLeftWith1InchMargin(t *testing.T) {

	milimeters := 25.4
	margin := 25.4

	output := SetAbsoluteHorizontalPrintPositionInMillimetersWithKnownMargin(uint32(milimeters), uint32(margin))

	if output[2] != 0 && output[3] != 0 {
		t.Errorf("Expected: 1, 104, Got: %d, %d", output[3], output[2])
	}

}
