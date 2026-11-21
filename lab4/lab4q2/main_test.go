package main

import "testing"

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %f", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("expected division by zero error")
	}
}

func TestSquareRoot(t *testing.T) {
	result, err := SquareRoot(25)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 5, got %f", result)
	}
}

func TestNegativeSquareRoot(t *testing.T) {
	_, err := SquareRoot(-9)
	if err == nil {
		t.Error("expected negative square root error")
	}
}

func TestValidateInput(t *testing.T) {
	result, err := ValidateInput("Hello Go")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if result != "Hello Go" {
		t.Errorf("expected Hello Go, got %s", result)
	}
}

func TestEmptyInput(t *testing.T) {
	_, err := ValidateInput("")
	if err == nil {
		t.Error("expected empty input error")
	}
}
