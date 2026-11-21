package main

import "fmt"

func calculateSquares(ch chan string) {
	fmt.Println("Square Goroutine started")

	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Square of %d = %d", i, i*i)
	}

	fmt.Println("Square Goroutine completed")
}

func calculateCubes(ch chan string) {
	fmt.Println("Cube Goroutine started")

	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Cube of %d = %d", i, i*i*i)
	}

	fmt.Println("Cube Goroutine completed")
}

func calculateFibonacci(ch chan string) {
	fmt.Println("Fibonacci Goroutine started")

	a, b := 0, 1

	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Fibonacci number %d = %d", i, a)
		a, b = b, a+b
	}

	fmt.Println("Fibonacci Goroutine completed")
}

func main() {
	fmt.Println("Main function started")

	ch := make(chan string, 15)

	go calculateSquares(ch)
	go calculateCubes(ch)
	go calculateFibonacci(ch)

	for i := 0; i < 15; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Main function completed")
}
