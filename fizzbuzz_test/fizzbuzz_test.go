package fizzbuzz_test

import (
	"golang-oop/fizzbuzz"
	"testing"
)

func TestFizzBuzz_1(t *testing.T) {

	// Act
	result := fizzbuzz.FizzBuzz(1)
	expected := "1"

	// Assert
	if result != expected {
		t.Errorf("FizzBuzz(1) result: %v : expected %v", result, expected)
	}
}
