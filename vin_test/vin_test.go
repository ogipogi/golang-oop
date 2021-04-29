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

const VIN_WL0 = "WL090001111224152"
const VIN_WA1 = "WA190001111224152"

/*

	1. Manufacturer first three digits

func TestVIN_WL0_from_VIN(t *testing.T) {
	// Arrange

	// Act
	manufacturer := vin.Manufacturer(VIN_WL0)

	// Assert
	if manufacturer != "WL0" {
		t.Errorf("unexpected manufactuer %s for VIN %s", manufacturer, VIN_WL0)
	}
}

func TestVIN_WA1_from_VIN(t *testing.T) {
	// Arrange

	// Act
	manufacturer := vin.Manufacturer(VIN_WA1)

	// Assert
	if manufacturer != "WA1" {
		t.Errorf("unexpected manufactuer %s for VIN %s", manufacturer, VIN_WA1)
	}
}
*/

/*

	1. Manufacturer first three digits
*/
func TestVIN_manufacturer_from_VIN(t *testing.T) {
	// Arrange
	vins := []struct {
		vin                  string
		expectedManufacturer string
	}{
		{VIN_WA1, "WA1"},
		{VIN_WL0, "WL0"},
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

/*
	if the last digit of the manufacturer ID is a 9
	the digits 12 to 14 are the second part of the ID

	VIN = W09000051T2123456
	MANUFACTURER_CODE = W09123
*/
