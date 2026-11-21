package main

import (
	"fmt"
	"math"
	"strings"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func SquareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("square root of a negative number is not allowed")
	}
	return math.Sqrt(n), nil
}

func ValidateInput(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", fmt.Errorf("input cannot be empty")
	}
	return input, nil
}

func main() {
	fmt.Println("Error Handling Demonstration")

	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", result)
	}

	_, err = Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	sqrt, err := SquareRoot(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square Root:", sqrt)
	}

	_, err = SquareRoot(-9)
	if err != nil {
		fmt.Println("Error:", err)
	}

	input, err := ValidateInput("")
	if err != nil {
		fmt.Println("Error:", err)
	}

	input, err = ValidateInput("Hello Go")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid Input:", input)
	}
}
