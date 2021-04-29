package vin_test

import (
	"golang-oop/vin"
	"testing"
)

/*
The unique vehicle identification number of every car includes—beside a
“running” (i.e., serial) number—information about the car, such as the
manufacturer, the producing factory, the car model, and if it is driven
from the left- or right-hand side.
*/


const VIN_WL0 = "WL090001111224152"
const VIN_WA1 = "WA190001111224152"

/*
	1. Manufacturer first three digits
 */
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
