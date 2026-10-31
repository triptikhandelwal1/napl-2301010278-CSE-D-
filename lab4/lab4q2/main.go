package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// Function to perform division
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

// Function to calculate square root
func SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("square root of a negative number is not allowed")
	}
	return math.Sqrt(n), nil
}

// Function to validate input
func ValidateInput(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", errors.New("input cannot be empty")
	}
	return input, nil
}

func main() {
	fmt.Println("Error Handling Demonstration")

	// Test division
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", result)
	}

	// Division by zero
	_, err = Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Test square root
	sqrt, err := SquareRoot(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square Root:", sqrt)
	}

	// Negative square root
	_, err = SquareRoot(-9)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Test empty input
	_, err = ValidateInput("")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Test valid input
	input, err := ValidateInput("Hello Go")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid Input:", input)
	}
}
