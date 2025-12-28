package main

import (
	"fmt"
)

func arithemeticOperation(op string) func(a, b int) int {
	// creating a switch that returns a function
	switch op {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "subtract":
		return func(a, b int) int {
			return a - b
		}
	case "multiply":
		return func(a, b int) int {
			return a * b
		}
	case "divide":
		return func(a, b int) int {
			if b == 0 {
				fmt.Println("division by zero is not possible")
			}
			return a / b
		}
	}
	return nil
}

func main() {
	var (
		op   string
		x, y int
	)
	fmt.Print("Enter any option between (add, subtract, multiply, divide) := ")
	fmt.Scan(&op)

	fmt.Print("Enter x :")
	fmt.Scan(&x)

	fmt.Print("Enter y : ")
	fmt.Scan(&y)
	// caling the function to perform operation
	function := arithemeticOperation(op)

	if function == nil {
		fmt.Println("Invalid operation!")
		return
	}
	answer := function(x, y)
	fmt.Printf("The result is: %d\n", answer)
}
