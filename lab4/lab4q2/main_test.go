package main

import (
	"testing"
)

// Test valid division
func TestDivideValid(t *testing.T) {
	result, err := Divide(10, 2)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

// Test division by zero
func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)

	if err == nil {
		t.Error("Expected division by zero error")
	}
}

// Test valid square root
func TestSquareRootValid(t *testing.T) {
	result, err := SquareRoot(25)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}

// Test negative square root
func TestSquareRootNegative(t *testing.T) {
	_, err := SquareRoot(-9)

	if err == nil {
		t.Error("Expected negative square root error")
	}
}

// Test valid input
func TestValidateInputValid(t *testing.T) {
	result, err := ValidateInput("Hello Go")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != "Hello Go" {
		t.Errorf("Expected Hello Go, got %v", result)
	}
}

// Test empty input
func TestValidateInputEmpty(t *testing.T) {
	_, err := ValidateInput("")

	if err == nil {
		t.Error("Expected empty input error")
	}
}
