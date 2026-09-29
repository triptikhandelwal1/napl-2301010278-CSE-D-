package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks int
}

func updateMarks(marks *int) {
	*marks = 90
}

func updateStudent(s1 *Student) {
	s1.Name = "Tripti"
	s1.Age = 20
	s1.Marks = 95
}

func main() {

	// Pointer manipulation
	marks := 75
	p := &marks

	// Default values
	fmt.Println("Default Marks:", marks)
	fmt.Println("Address of Marks:", &marks)
	fmt.Println("Value using pointer:", *p)

	// Function call
	updateMarks(&marks)

	// Updated marks
	fmt.Println("Marks after function call:", marks)

	// Create structure using new()
	student := new(Student)

	// Default student values
	fmt.Println("\nDefault Student Values:")
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Marks:", student.Marks)

	// Function call
	updateStudent(student)

	// Updated student values
	fmt.Println("\nUpdated Student Values:")
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Marks:", student.Marks)
}
