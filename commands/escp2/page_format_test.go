package escp2

import "testing"

func TestSetPageFormatA4(t *testing.T) {
	// When using 1/360 as a unit page size shall be 15 * 256 + 113

	lower := byte(113)
	upper := byte(16)

	payload := SetPageFormatA4()
	if !(payload[9] == lower && payload[10] == upper) {
		t.Errorf("SetPageFormatA4: expected 113 and 16, got %d and %d", payload[9], payload[10])
	}

}
