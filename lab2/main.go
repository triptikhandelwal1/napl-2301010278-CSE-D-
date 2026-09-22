package main

import (
	"fmt"
	"lab2/mathutil"
	"lab2/strop"
)

func main() {

	var n int
	var a, b int
	var text string

	fmt.Print("Enter a number for factorial: ")
	fmt.Scan(&n)

	fmt.Print("Enter base: ")
	fmt.Scan(&a)

	fmt.Print("Enter exponent: ")
	fmt.Scan(&b)

	fmt.Print("Enter a string: ")
	fmt.Scan(&text)

	fmt.Println()
	fmt.Println("Factorial:", mathutil.Fact(n))
	fmt.Println("Power:", mathutil.Pow(a, b))
	fmt.Println("Reversed string:", strop.Reverse(text))
	fmt.Println("Number of vowels:", strop.CountVowels(text))
}
