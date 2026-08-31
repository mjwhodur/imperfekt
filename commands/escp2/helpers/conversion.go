package helpers

import (
	"errors"
)

func MilimeterToUnit(milimeter uint32, unit uint32) (uint32, error) {

	if (unit == 5) || (unit == 10) || (unit == 20) || (unit == 30) || (unit == 40) || (unit == 50) || (unit == 60) {
		inInches := float32(milimeter) / 25.4

		inUnits := inInches * 3600 / float32(unit)

		return uint32(inUnits), nil

	} else {
		return 0, errors.New("milimeterToUnit called with invalid arguments")
	}

}
