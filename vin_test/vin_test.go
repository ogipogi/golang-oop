package vin_test

import (
	"golang-oop/vin"
	"testing"
)

/*
The unique vehicle identification number of every car includes—beside a
“running” (i.e., serial) number—information about the car, such as the
expectedManufacturer, the producing factory, the car model, and if it is driven
from the left- or right-hand side.
*/

/*

	1. Manufacturer first three digits
	2. if the last digit of the manufacturer ID is a 9
	the digits 11 to 14 are the second part of the ID

	VIN = W09000051T2123456
	MANUFACTURER_CODE = W09123
*/
func TestVIN_manufacturer_from_VIN(t *testing.T) {
	// Arrange
	vins := []struct {
		vin                  string
		expectedManufacturer string
	}{
		{"WA190001111224152", "WA1"},
		{"WL090001111224152", "WL0"},
		{"W09000051T2123456", "W09123"},
		{"AB900011112241528", "AB9241"},
	}

	for _, number := range vins {
		// Act
		manufacturer := vin.Manufacturer(number.vin)

		// Assert
		if manufacturer != number.expectedManufacturer {
			t.Errorf("unexpected manufactuer %s for VIN %s", manufacturer, number.vin)
		}
	}

}

