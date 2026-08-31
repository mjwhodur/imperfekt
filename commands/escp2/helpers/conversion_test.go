package helpers

import "testing"

func TestConversion1inch(t *testing.T) {
	milimeters := 254

	shallBe360, err := MilimeterToUnit(uint32(milimeters), uint32(10))
	if err != nil {
		t.Error(err)
	}
	if shallBe360 != 3600 {
		t.Errorf("254 millimeters (10 inch) shall be 3600 * 1 / 360 printer units, was %d", shallBe360)
	}
}

func TestConversion5cm(t *testing.T) {
	milimeters := 50

	shallBe709, err := MilimeterToUnit(uint32(milimeters), uint32(10))
	if err != nil {
		t.Error(err)
	}
	if shallBe709 != 708 {
		t.Errorf("500 millimeters shall be 360 * 1 / 360 printer units, was %d", shallBe709)
	}
}

func TestConversionA4(t *testing.T) {
	milimeters := 297

	shallBe709, err := MilimeterToUnit(uint32(milimeters), uint32(10))
	if err != nil {
		t.Error(err)
	}
	if shallBe709 != 4209 {
		t.Errorf("254 millimeters shall be 360 * 1 / 360 printer units, was %d", shallBe709)
	}
}
